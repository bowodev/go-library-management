package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/bowodev/go-library-management/internal/utils"
)

type createBook struct {
	repository interfaces.IRepository
	cacher     interfaces.ICache[dto.Book]
	validator  interfaces.IValidator
}

var _ interfaces.ICreateBook = (*createBook)(nil)

func NewCreateBook(
	repository interfaces.IRepository,
	cacher interfaces.ICache[dto.Book],
	validator interfaces.IValidator,
) *createBook {
	return &createBook{
		repository: repository,
		cacher:     cacher,
		validator:  validator,
	}
}

// Do implements interfaces.ICreateBook.
func (b *createBook) Do(ctx context.Context, in dto.Book) (dto.Book, error) {
	if in.Author.ID < 1 {
		return dto.Book{}, utils.ErrInvalidAuthorId
	}

	if err := b.validator.Do(ctx, in); err != nil {
		return dto.Book{}, err
	}

	trx, err := b.repository.Begin(ctx)
	if err != nil {
		return dto.Book{}, err
	}

	book, err := b.repository.CreateBook(ctx, trx, in)
	if err != nil {
		log.Println("error happen when create book, err: ", err)
		if err := trx.Rollback(ctx); err != nil {
			return dto.Book{}, err
		}
		return dto.Book{}, err
	}

	if err := trx.Commit(ctx); err != nil {
		if err := trx.Rollback(ctx); err != nil {
			return dto.Book{}, err
		}
		return dto.Book{}, err
	}

	if err := b.cacher.Set(ctx, fmt.Sprintf("%d", book.ID), book); err != nil {
		log.Println("error happen when try to set cache, err: ", err)
	}

	return book, nil
}
