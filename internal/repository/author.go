package repository

import (
	"context"

	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
)

// CreateAuthor implements interfaces.IRepository.
func (r *repository) CreateAuthor(ctx context.Context, tx interfaces.ITransaction, in dto.Author) (dto.Author, error) {
	panic("unimplemented")
}

// DeleteAuthor implements interfaces.IRepository.
func (r *repository) DeleteAuthor(ctx context.Context, tx interfaces.ITransaction, id int64) error {
	panic("unimplemented")
}

// GetAuthorAll implements interfaces.IRepository.
func (r *repository) GetAuthorAll(ctx context.Context, tx interfaces.ITransaction) (dto.Authors, error) {
	panic("unimplemented")
}

// GetAuthorById implements interfaces.IRepository.
func (r *repository) GetAuthorById(ctx context.Context, tx interfaces.ITransaction, id int64) (dto.Author, error) {
	panic("unimplemented")
}

// UpdateAuthor implements interfaces.IRepository.
func (r *repository) UpdateAuthor(ctx context.Context, tx interfaces.ITransaction, in dto.Author) (dto.Author, error) {
	panic("unimplemented")
}
