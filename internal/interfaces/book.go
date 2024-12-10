package interfaces

import (
	"context"

	"github.com/bowodev/go-library-management/internal/dto"
)

//go:generate mockgen -source=book.go -destination=./mock/book_mock.go -package=mocks
type ICreateBook interface {
	Do(ctx context.Context, in dto.Book) (dto.Book, error)
}
