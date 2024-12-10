package usecase_test

import (
	"github.com/bowodev/go-library-management/internal/dto"
	mocks "github.com/bowodev/go-library-management/internal/interfaces/mock"
	"github.com/stretchr/testify/assert"
)

type CreateBookTestCase struct {
	name           string
	createBookArgs dto.Book
	mockFunc       func(md *CreateBookDeps)
	wantErr        assert.ErrorAssertionFunc
}

type CreateBookDeps struct {
	repository *mocks.MockIRepository
	trx        *mocks.MockITransaction
	validator  *mocks.MockIValidator
	cache      *mocks.MockICache[dto.Book]
}
