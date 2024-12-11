package repository

import (
	"context"
	"errors"
	"time"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/bowodev/go-library-management/internal/model"
	"github.com/bowodev/go-library-management/internal/utils"
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
		return dto.Book{}, utils.ErrDuplicateBookTitle
	}

	if err := r.getDb(tx).Table("books").Create(&book).Error; err != nil {
		return dto.Book{}, err
	}

	in.ID = book.ID
	return in, nil
}

// DeleteBook implements interfaces.IRepository.
func (r *repository) DeleteBook(ctx context.Context, tx interfaces.ITransaction, id int64) error {
	return r.getDb(tx).Table("books").
		Where(`deleted_at isnull and id = ?`, id).
		UpdateColumn("deleted_at", time.Now()).
		Error
}

// GetBookAll implements interfaces.IRepository.
func (r *repository) GetBookAll(ctx context.Context, tx interfaces.ITransaction) (dto.Books, error) {
	var (
		books   []model.Book
		authors []model.Author

		res = dto.Books{}
	)

	err := r.getDb(tx).Table("books").
		Where(`deleted_at isnull`).
		Scan(&books).
		Error
	if err != nil {
		return dto.Books{}, err
	}

	authorIds := []int64{}
	for _, v := range books {
		authorIds = append(authorIds, v.ID)
	}

	err = r.getDb(tx).Table("authors").
		Where(`deleted_at isnull and id in (?)`, authorIds).
		Scan(&authors).
		Error
	if err != nil {
		return dto.Books{}, err
	}

	mapOfAuthors := map[int64]model.Author{}
	for _, a := range authors {
		mapOfAuthors[a.ID] = a
	}

	for idx, b := range books {
		if _, ok := mapOfAuthors[b.AuthorId]; ok {
			books[idx].Author = mapOfAuthors[b.AuthorId]
		}
	}

	res.FromModels(books)
	return res, nil
}

// GetBookByAuthor implements interfaces.IRepository.
func (r *repository) GetBookByAuthor(ctx context.Context, tx interfaces.ITransaction, authorId int64) (dto.Books, error) {
	var (
		books  []model.Book
		author model.Author

		res = dto.Books{}
	)

	err := r.getDb(tx).Table("authors").
		Where(`deleted_at isnull and id = ?`, authorId).
		Scan(&author).
		Error
	if err != nil {
		return dto.Books{}, err
	}

	err = r.getDb(tx).Table("books").
		Where(`deleted_at isnull and author_id = ?`, authorId).
		Scan(&books).
		Error
	if err != nil {
		return dto.Books{}, err
	}

	for idx := range books {
		books[idx].Author = author
	}

	res.FromModels(books)

	return res, nil
}

// GetBookById implements interfaces.IRepository.
func (r *repository) GetBookById(ctx context.Context, tx interfaces.ITransaction, id int64) (dto.Book, error) {
	var (
		book   model.Book
		author model.Author

		res = dto.Book{}
	)

	err := r.getDb(tx).Table("books").
		Where(`deleted_at isnull and id = ?`, id).
		Scan(&book).
		Error
	if err != nil {
		return dto.Book{}, err
	}

	err = r.getDb(tx).Table("authors").
		Where(`deleted_at isnull and id = ?`, book.AuthorId).
		Scan(&author).
		Error
	if err != nil {
		return dto.Book{}, err
	}

	book.Author = author

	res.FromModel(book)

	return res, nil
}

// UpdateBook implements interfaces.IRepository.
func (r *repository) UpdateBook(ctx context.Context, tx interfaces.ITransaction, in dto.Book) (dto.Book, error) {
	var (
		checkAuthor model.Author
		checkBook   model.Book
		payload     = in.ToModel()
	)

	err := r.getDb(tx).
		Table("authors").
		Where(`deleted_at isnull and id = ?`, in.Author.ID).
		First(&checkAuthor).
		Error

	if err != nil {
		return dto.Book{}, err
	}
	if checkAuthor.ID < 1 {
		return dto.Book{}, utils.ErrInvalidAuthorId
	}

	err = r.getDb(tx).
		Table("books").
		Where(`deleted_at isnull and title = ?`, in.Title).
		First(&checkBook).
		Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.Book{}, err
	}
	if checkBook.ID > 0 && checkBook.ID != in.ID {
		return dto.Book{}, utils.ErrDuplicateBookTitle
	}

	err = r.getDb(tx).
		Table("books").
		Updates(&payload).
		Where(`deleted_at isnull and id = ?`, in.ID).
		Error
	if err != nil {
		return dto.Book{}, err
	}

	return in, nil
}
