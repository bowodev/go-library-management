package usecase

import (
	"context"
	"fmt"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/gofiber/fiber/v2/log"
)

type getBookByAuthor struct {
	repository interfaces.IRepository
	cacher     interfaces.ICache[dto.Books]
}

// Do implements interfaces.IGetBookByAuthor.
func (g *getBookByAuthor) Do(ctx context.Context, authorId int64) (dto.Books, error) {
	books, _ := g.cacher.Get(ctx, fmt.Sprintf("%d", authorId))
	if len(books) > 0 {
		return books, nil
	}

	books, err := g.repository.GetBookByAuthor(ctx, nil, authorId)
	if err != nil {
		log.Error("failed to get book, error", err)
		return books, err
	}

	if err := g.cacher.Set(ctx, fmt.Sprintf("%d", authorId), books); err != nil {
		log.Warn("failed to set cache, error", err)
	}

	return books, nil
}

var _ interfaces.IGetBookByAuthor = (*getBookByAuthor)(nil)

func NewGetBookByAuthor(
	repository interfaces.IRepository,
	cacher interfaces.ICache[dto.Books],
) *getBookByAuthor {
	return &getBookByAuthor{
		repository: repository,
		cacher:     cacher,
	}
}
