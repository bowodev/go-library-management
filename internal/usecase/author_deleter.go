package usecase

import (
	"context"
	"fmt"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/gofiber/fiber/v2/log"
)

type deleteAuthor struct {
	repository   interfaces.IRepository
	authorCacher interfaces.ICache[dto.Author]
	bookCacher   interfaces.ICache[dto.Book]
}

var _ interfaces.IDeleteAuthor = (*deleteAuthor)(nil)

func NewAuthorDeleter(
	repository interfaces.IRepository,
	authorCacher interfaces.ICache[dto.Author],
	bookCacher interfaces.ICache[dto.Book],
) *deleteAuthor {
	return &deleteAuthor{
		repository:   repository,
		authorCacher: authorCacher,
		bookCacher:   bookCacher,
	}
}

// Do implements interfaces.IDeleteAuthor.
func (d *deleteAuthor) Do(ctx context.Context, id int64) error {
	books, err := d.repository.GetBookByAuthor(ctx, nil, id)
	if err != nil {
		return err
	}

	tx, err := d.repository.Begin(ctx)
	if err != nil {
		return err
	}

	if err := d.repository.DeleteAuthor(ctx, tx, id); err != nil {
		log.Error("failed to delete author", err)
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

	if err := d.authorCacher.Del(ctx, fmt.Sprintf("%d", id)); err != nil {
		log.Warn("failed to delete author cache", err)
	}

	for _, b := range books {
		if err := d.bookCacher.Del(ctx, fmt.Sprintf("%d", b.ID)); err != nil {
			log.Warn("failed to delete book cache", err)
		}
	}

	return nil
}
