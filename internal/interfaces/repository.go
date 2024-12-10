package interfaces

import (
	"context"

	"github.com/bowodev/go-library-management/internal/dto"
)

//go:generate mockgen -source=repository.go -destination=./mock/repository_mock.go -package=mocks
type ITransaction interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type IRepository interface {
	Begin(ctx context.Context) (ITransaction, error)

	IAuthorRepository
	IBookRepository
}

type IAuthorRepository interface {
	CreateAuthor(ctx context.Context, tx ITransaction, in dto.Author) (dto.Author, error)
	GetAuthorById(ctx context.Context, tx ITransaction, id int64) (dto.Author, error)
	GetAuthorAll(ctx context.Context, tx ITransaction) (dto.Authors, error)
	UpdateAuthor(ctx context.Context, tx ITransaction, in dto.Author) (dto.Author, error)
	DeleteAuthor(ctx context.Context, tx ITransaction, id int64) error
}

type IBookRepository interface {
	CreateBook(ctx context.Context, tx ITransaction, in dto.Book) (dto.Book, error)
	GetBookById(ctx context.Context, tx ITransaction, id int64) (dto.Book, error)
	GetBookByAuthor(ctx context.Context, tx ITransaction, authorId int64) (dto.Books, error)
	GetBookAll(ctx context.Context, tx ITransaction) (dto.Books, error)
	UpdateBook(ctx context.Context, tx ITransaction, in dto.Book) (dto.Book, error)
	DeleteBook(ctx context.Context, tx ITransaction, id int64) error
}
