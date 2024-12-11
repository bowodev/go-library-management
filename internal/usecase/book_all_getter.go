package usecase

import (
	"context"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
)

type getAllBook struct {
	repository interfaces.IRepository
}

var _ interfaces.IGetBookAll = (*getAllBook)(nil)

func NewBookAllGetter(
	repository interfaces.IRepository,
) *getAllBook {
	return &getAllBook{
		repository: repository,
	}
}

// Do implements interfaces.IGetBookAll.
func (g *getAllBook) Do(ctx context.Context) (dto.Books, error) {
	books, err := g.repository.GetBookAll(ctx, nil)
	if err != nil {
		return dto.Books{}, err
	}

	return books, nil
}
