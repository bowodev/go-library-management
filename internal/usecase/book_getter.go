package usecase

import (
	"context"
	"fmt"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/gofiber/fiber/v2/log"
)

type getBook struct {
	cache      interfaces.ICache[dto.Book]
	repository interfaces.IRepository
}

// Do implements interfaces.IGetBookById.
func (g *getBook) Do(ctx context.Context, id int64) (dto.Book, error) {
	book, _ := g.cache.Get(ctx, fmt.Sprintf("%d", id))
	if book.ID > 0 {
		return book, nil
	}

	book, err := g.repository.GetBookById(ctx, nil, id)
	if err != nil {
		log.Error("failed to get book, error", err)
		return book, err
	}

	if err := g.cache.Set(ctx, fmt.Sprintf("%d", id), book); err != nil {
		log.Warn("failed to set cache", err)
	}

	return book, nil
}

func NewGetBook(
	cache interfaces.ICache[dto.Book],
	repository interfaces.IRepository,
) *getBook {
	return &getBook{
		cache:      cache,
		repository: repository,
	}
}

var _ interfaces.IGetBookById = (*getBook)(nil)
