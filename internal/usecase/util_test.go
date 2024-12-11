package usecase_test

import (
	"github.com/bowodev/go-library-management/internal/dto"
	mocks "github.com/bowodev/go-library-management/internal/interfaces/mock"
	"github.com/stretchr/testify/assert"
)

// CreateAuthor...
type CreateAuthorTestCase struct {
	name             string
	createAuthorArgs dto.Author
	mockFunc         func(md *CreateAuthorDeps)
	wantErr          assert.ErrorAssertionFunc
}

type CreateAuthorDeps struct {
	repository *mocks.MockIRepository
	trx        *mocks.MockITransaction
	validator  *mocks.MockIValidator
	cache      *mocks.MockICache[dto.Author]
}

// GetAuthorAll...
type GetAuthorAllTestCase struct {
	name     string
	mockFunc func(md *GetAuthorAllDeps)
	wantErr  assert.ErrorAssertionFunc
}

type GetAuthorAllDeps struct {
	repository *mocks.MockIRepository
	cache      *mocks.MockICache[dto.Authors]
}

// GetAuthorById...
type GetAuthorByIdTestCase struct {
	name              string
	getAuthorByIdArgs int64
	mockFunc          func(md *GetAuthorByIdDeps)
	wantErr           assert.ErrorAssertionFunc
}

type GetAuthorByIdDeps struct {
	repository *mocks.MockIRepository
	cache      *mocks.MockICache[dto.Author]
}

// UpdateAuthor
type UpdateAuthorTestCase struct {
	name             string
	updateAuthorArgs dto.Author
	mockFunc         func(md *UpdateAuthorDeps)
	wantErr          assert.ErrorAssertionFunc
}

type UpdateAuthorDeps struct {
	repository *mocks.MockIRepository
	trx        *mocks.MockITransaction
	validator  *mocks.MockIValidator
	cache      *mocks.MockICache[dto.Author]
}

// DeleteAuthor
type DeleteAuthorTestCase struct {
	name             string
	deleteAuthorArgs int64
	mockFunc         func(md *DeleteAuthorDeps)
	wantErr          assert.ErrorAssertionFunc
}

type DeleteAuthorDeps struct {
	repository  *mocks.MockIRepository
	trx         *mocks.MockITransaction
	authorCache *mocks.MockICache[dto.Author]
	bookCache   *mocks.MockICache[dto.Book]
}

// CreateBook...
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
