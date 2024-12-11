package repository

import (
	"context"
	"time"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/bowodev/go-library-management/internal/model"
)

// CreateAuthor implements interfaces.IRepository.
func (r *repository) CreateAuthor(ctx context.Context, tx interfaces.ITransaction, in dto.Author) (dto.Author, error) {
	author := in.ToModel()
	if err := r.getDb(tx).Table("authors").Create(&author).Error; err != nil {
		return dto.Author{}, err
	}

	in.ID = author.ID
	return in, nil
}

// DeleteAuthor implements interfaces.IRepository.
func (r *repository) DeleteAuthor(ctx context.Context, tx interfaces.ITransaction, id int64) error {
	var check model.Author

	err := r.getDb(tx).Table("authors").
		Where(`id = ? and deleted_at isnull`, id).
		First(&check).
		Error
	if err != nil {
		return err
	}

	return r.getDb(tx).Table("authors").
		Where(`id = ? and deleted_at isnull`, id).
		UpdateColumns(map[string]any{
			"deleted_at": time.Now(),
		}).
		Error
}

// GetAuthorAll implements interfaces.IRepository.
func (r *repository) GetAuthorAll(ctx context.Context, tx interfaces.ITransaction) (dto.Authors, error) {
	var (
		authors []model.Author
		res     dto.Authors
	)

	err := r.getDb(tx).Table("authors").
		Where(`deleted_at isnull`).
		Scan(&authors).
		Error
	if err != nil {
		return dto.Authors{}, err
	}

	res.FromModels(authors)
	return res, nil
}

// GetAuthorById implements interfaces.IRepository.
func (r *repository) GetAuthorById(ctx context.Context, tx interfaces.ITransaction, id int64) (dto.Author, error) {
	var (
		author model.Author
		res    dto.Author
	)

	err := r.getDb(tx).Table("authors").
		Where(`id = ? and deleted_at isnull`, id).
		First(&author).
		Error
	if err != nil {
		return dto.Author{}, err
	}

	res.FromModel(author)
	return res, nil
}

// UpdateAuthor implements interfaces.IRepository.
func (r *repository) UpdateAuthor(ctx context.Context, tx interfaces.ITransaction, in dto.Author) (dto.Author, error) {
	var check model.Author

	err := r.getDb(tx).Table("authors").
		Where(`id = ? and deleted_at isnull`, in.ID).
		First(&check).
		Error
	if err != nil {
		return dto.Author{}, err
	}

	updatedAuthor := in.ToModel()
	err = r.getDb(tx).Table("authors").
		Where(`id = ? and deleted_at isnull`, in.ID).
		Updates(&updatedAuthor).
		Error
	if err != nil {
		return dto.Author{}, err
	}

	return in, nil
}
