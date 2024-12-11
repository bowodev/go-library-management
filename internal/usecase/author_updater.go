package usecase

import (
	"context"
	"fmt"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/gofiber/fiber/v2/log"
)

type updateAuthor struct {
	repository interfaces.IRepository
	cacher     interfaces.ICache[dto.Author]
	validator  interfaces.IValidator
}

var _ interfaces.IUpdateAuthor = (*updateAuthor)(nil)

func NewAuthorUpdater(
	repository interfaces.IRepository,
	cacher interfaces.ICache[dto.Author],
	validator interfaces.IValidator,
) *updateAuthor {
	return &updateAuthor{
		repository: repository,
		cacher:     cacher,
		validator:  validator,
	}
}

// Do implements interfaces.IUpdateAuthor.
func (u *updateAuthor) Do(ctx context.Context, in dto.Author) (dto.Author, error) {
	if err := u.validator.Do(ctx, in); err != nil {
		return dto.Author{}, err
	}

	tx, err := u.repository.Begin(ctx)
	if err != nil {
		return dto.Author{}, err
	}

	author, err := u.repository.UpdateAuthor(ctx, tx, in)
	if err != nil {
		log.Error("failed to update author, error:", err)
		if err := tx.Rollback(ctx); err != nil {
			return dto.Author{}, err
		}
		return dto.Author{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return dto.Author{}, err
		}
		return dto.Author{}, err
	}

	if err := u.cacher.Set(ctx, fmt.Sprintf("%d", author.ID), author); err != nil {
		log.Warn("failed to set author cache, error:", err)
	}

	return author, nil
}
