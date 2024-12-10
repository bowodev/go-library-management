package repository

import (
	"context"
	"errors"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/bowodev/go-library-management/internal/model"
	"gorm.io/gorm"
)

// CreateBook implements interfaces.IRepository.
func (r *repository) CreateBook(ctx context.Context, tx interfaces.ITransaction, in dto.Book) (dto.Book, error) {
	book := in.ToModel()

	var check model.Book
	err := r.getDb(tx).Table("books").
		Where(`"title" = ? and deleted_at isnull`, in.Title).
		First(&check).
		Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.Book{}, err
	}

	if check.ID > 0 {
		return dto.Book{}, errors.New("duplicate book name")
	}

	if err := r.getDb(tx).Table("books").Create(&book).Error; err != nil {
		return dto.Book{}, err
	}

	in.ID = book.ID
	return in, nil
}

// DeleteBook implements interfaces.IRepository.
func (r *repository) DeleteBook(ctx context.Context, tx interfaces.ITransaction, id int64) error {
	panic("unimplemented")
}

// GetBookAll implements interfaces.IRepository.
func (r *repository) GetBookAll(ctx context.Context, tx interfaces.ITransaction) (dto.Books, error) {
	panic("unimplemented")
}

// GetBookByAuthor implements interfaces.IRepository.
func (r *repository) GetBookByAuthor(ctx context.Context, tx interfaces.ITransaction, authorId int64) (dto.Books, error) {
	panic("unimplemented")
}

// GetBookById implements interfaces.IRepository.
func (r *repository) GetBookById(ctx context.Context, tx interfaces.ITransaction, id int64) (dto.Book, error) {
	panic("unimplemented")
}

// UpdateBook implements interfaces.IRepository.
func (r *repository) UpdateBook(ctx context.Context, tx interfaces.ITransaction, in dto.Book) (dto.Book, error) {
	panic("unimplemented")
}
