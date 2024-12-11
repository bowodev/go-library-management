package interfaces

import (
	"context"

	"github.com/bowodev/go-library-management/internal/dto"
)

type ICreateAuthor interface {
	Do(context.Context, dto.Author) (dto.Author, error)
}

type IGetAuthorById interface {
	Do(context.Context, int64) (dto.Author, error)
}

type IGetAuthorAll interface {
	Do(context.Context) (dto.Authors, error)
}

type IUpdateAuthor interface {
	Do(context.Context, dto.Author) (dto.Author, error)
}

type IDeleteAuthor interface {
	Do(context.Context, int64) error
}
