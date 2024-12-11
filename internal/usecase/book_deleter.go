package usecase

import (
	"context"
	"fmt"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/gofiber/fiber/v2/log"
)

type deleteBook struct {
	repository interfaces.IRepository
	cacher     interfaces.ICache[dto.Book]
}

// Do implements interfaces.IDeleteBook.
func (d *deleteBook) Do(ctx context.Context, id int64) error {
	tx, err := d.repository.Begin(ctx)
	if err != nil {
		return err
	}

	if err := d.repository.DeleteBook(ctx, tx, id); err != nil {
		log.Error("failed to delete book, error", err)
		if err := tx.Rollback(ctx); err != nil {
			return err
		}

		return err
	}

	if err := tx.Commit(ctx); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}

		return err
	}

	if err := d.cacher.Del(ctx, fmt.Sprintf("%d", id)); err != nil {
		log.Warn("failed to del cache, error", err)
	}
	return nil
}

func NewDeleteBook(
	repository interfaces.IRepository,
	cacher interfaces.ICache[dto.Book],
) *deleteBook {
	return &deleteBook{
		repository: repository,
		cacher:     cacher,
	}
}

var _ interfaces.IDeleteBook = (*deleteBook)(nil)
