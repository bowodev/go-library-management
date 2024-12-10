package repository

import (
	"context"

	"github.com/bowodev/go-library-management/internal/interfaces"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type transaction struct {
	tx *gorm.DB
}

var _ interfaces.IRepository = (*repository)(nil)
var _ interfaces.ITransaction = (*transaction)(nil)

func New(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Begin(ctx context.Context) (interfaces.ITransaction, error) {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &transaction{
		tx: tx,
	}, nil
}

func (t *transaction) Commit(ctx context.Context) error {
	return t.tx.WithContext(ctx).Commit().Error
}

func (t *transaction) Rollback(ctx context.Context) error {
	return t.tx.WithContext(ctx).Rollback().Error
}

func (r *repository) getDb(tx interfaces.ITransaction) *gorm.DB {
	var db *gorm.DB

	if tx != nil {
		db = tx.(*transaction).tx
	} else {
		db = r.db
	}

	return db
}
