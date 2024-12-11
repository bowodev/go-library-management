package usecase

import (
	"context"
	"fmt"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/gofiber/fiber/v2/log"
)

type getAuthor struct {
	repository interfaces.IRepository
	cacher     interfaces.ICache[dto.Author]
}

var _ interfaces.IGetAuthorById = (*getAuthor)(nil)

func NewAuthorGetter(
	repository interfaces.IRepository,
	cacher interfaces.ICache[dto.Author],
) *getAuthor {
	return &getAuthor{
		repository: repository,
		cacher:     cacher,
	}
}

// Do implements interfaces.IGetAuthorById.
func (g *getAuthor) Do(ctx context.Context, id int64) (dto.Author, error) {
	author, _ := g.cacher.Get(ctx, fmt.Sprintf("%d", id))
	if author.ID > 0 {
		return author, nil
	}

	author, err := g.repository.GetAuthorById(ctx, nil, id)
	if err != nil {
		log.Error("failed to get author", err)
		return author, err
	}

	if err := g.cacher.Set(ctx, fmt.Sprintf("%d", id), author); err != nil {
		log.Warn("failed to set cache", err)
	}

	return author, nil
}
