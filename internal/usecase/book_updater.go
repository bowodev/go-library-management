package usecase

import (
	"context"
	"fmt"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/gofiber/fiber/v2/log"
)

type updateBook struct {
	repository interfaces.IRepository
	validator  interfaces.IValidator
	cacher     interfaces.ICache[dto.Book]
}

// Do implements interfaces.IUpdateBook.
func (u *updateBook) Do(ctx context.Context, in dto.Book) (dto.Book, error) {
	if err := u.validator.Do(ctx, in); err != nil {
		return dto.Book{}, err
	}

	tx, err := u.repository.Begin(ctx)
	if err != nil {
		return dto.Book{}, err
	}

	book, err := u.repository.UpdateBook(ctx, tx, in)
	if err != nil {
		log.Error("failed to update book, error", err)
		if err := tx.Rollback(ctx); err != nil {
			return dto.Book{}, err
		}

		return dto.Book{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return dto.Book{}, err
		}

		return dto.Book{}, err
	}

	if err := u.cacher.Set(ctx, fmt.Sprintf("%d", in.ID), book); err != nil {
		log.Warn("failed to set cache", err)
	}

	return book, nil
}

func NewUpdateBook(
	repository interfaces.IRepository,
	validator interfaces.IValidator,
	cacher interfaces.ICache[dto.Book],
) *updateBook {
	return &updateBook{
		repository: repository,
		validator:  validator,
		cacher:     cacher,
	}
}

var _ interfaces.IUpdateBook = (*updateBook)(nil)
