package usecase

import (
	"context"
	"fmt"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/gofiber/fiber/v2/log"
)

type createAuthor struct {
	repository interfaces.IRepository
	cacher     interfaces.ICache[dto.Author]
	validator  interfaces.IValidator
}

var _ interfaces.ICreateAuthor = (*createAuthor)(nil)

func NewAuthorCreator(
	repository interfaces.IRepository,
	cacher interfaces.ICache[dto.Author],
	validator interfaces.IValidator,
) *createAuthor {
	return &createAuthor{
		repository: repository,
		cacher:     cacher,
		validator:  validator,
	}
}

// Do implements interfaces.ICreateAuthor.
func (c *createAuthor) Do(ctx context.Context, in dto.Author) (dto.Author, error) {
	if err := c.validator.Do(ctx, in); err != nil {
		return dto.Author{}, err
	}

	tx, err := c.repository.Begin(ctx)
	if err != nil {
		return dto.Author{}, err
	}

	author, err := c.repository.CreateAuthor(ctx, tx, in)
	if err != nil {
		log.Error("failed to create author", err)
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

	if err := c.cacher.Set(ctx, fmt.Sprintf("%d", author.ID), author); err != nil {
		log.Warn("failed to set cache")
	}

	return author, nil
}
