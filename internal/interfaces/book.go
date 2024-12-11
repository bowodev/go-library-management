package interfaces

import (
	"context"

	"github.com/bowodev/go-library-management/internal/dto"
)

//go:generate mockgen -source=book.go -destination=./mock/book_mock.go -package=mocks
type (
	ICreateBook interface {
		Do(ctx context.Context, in dto.Book) (dto.Book, error)
	}

	IGetBookById interface {
		Do(ctx context.Context, id int64) (dto.Book, error)
	}

	IGetBookAll interface {
		Do(ctx context.Context) (dto.Books, error)
	}

	IGetBookByAuthor interface {
		Do(ctx context.Context, authorId int64) (dto.Books, error)
	}

	IUpdateBook interface {
		Do(ctx context.Context, in dto.Book) (dto.Book, error)
	}

	IDeleteBook interface {
		Do(ctx context.Context, id int64) error
	}
)
