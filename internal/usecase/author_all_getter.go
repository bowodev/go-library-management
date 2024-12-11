package usecase

import (
	"context"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
)

type getAuthorAll struct {
	repository interfaces.IRepository
}

func NewAuthorAllGetter(
	repository interfaces.IRepository,
) *getAuthorAll {
	return &getAuthorAll{
		repository: repository,
	}
}

var _ interfaces.IGetAuthorAll = (*getAuthorAll)(nil)

// Do implements interfaces.IGetAuthorAll.
func (g *getAuthorAll) Do(ctx context.Context) (dto.Authors, error) {
	authors, err := g.repository.GetAuthorAll(ctx, nil)
	if err != nil {
		return authors, err
	}

	return authors, nil
}
