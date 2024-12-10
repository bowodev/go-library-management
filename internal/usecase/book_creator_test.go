package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/bowodev/go-library-management/internal/dto"
	mocks "github.com/bowodev/go-library-management/internal/interfaces/mock"
	"github.com/bowodev/go-library-management/internal/usecase"
	"github.com/bowodev/go-library-management/internal/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateBook_HappyCase(t *testing.T) {
	testCases := []CreateBookTestCase{
		{
			name: "success",
			createBookArgs: dto.Book{
				Title:       "some book",
				Description: "description",
				PublishDate: time.Now(),
				Author: dto.Author{
					ID: 1,
				},
			},
			mockFunc: func(md *CreateBookDeps) {
				md.validator.EXPECT().Do(gomock.Any(), gomock.Any()).Return(nil)
				md.repository.EXPECT().Begin(gomock.Any()).Return(md.trx, nil)
				md.repository.EXPECT().CreateBook(gomock.Any(), gomock.Any(), gomock.Any()).Return(dto.Book{
					ID:          1,
					Title:       "some book",
					Description: "description",
					PublishDate: time.Now(),
					Author: dto.Author{
						ID: 1,
					},
				}, nil)
				md.trx.EXPECT().Commit(gomock.Any()).Return(nil)
				md.cache.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
		{
			name: "set cache error, expected return no error",
			createBookArgs: dto.Book{
				Title:       "some book",
				Description: "description",
				PublishDate: time.Now(),
				Author: dto.Author{
					ID: 1,
				},
			},
			mockFunc: func(md *CreateBookDeps) {
				md.validator.EXPECT().Do(gomock.Any(), gomock.Any()).Return(nil)
				md.repository.EXPECT().Begin(gomock.Any()).Return(md.trx, nil)
				md.repository.EXPECT().CreateBook(gomock.Any(), gomock.Any(), gomock.Any()).Return(dto.Book{
					ID:          1,
					Title:       "some book",
					Description: "description",
					PublishDate: time.Now(),
					Author: dto.Author{
						ID: 1,
					},
				}, nil)
				md.trx.EXPECT().Commit(gomock.Any()).Return(nil)
				md.cache.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			repository := mocks.NewMockIRepository(c)
			trx := mocks.NewMockITransaction(c)
			validator := mocks.NewMockIValidator(c)
			cache := mocks.NewMockICache[dto.Book](c)

			if tt.mockFunc != nil {
				tt.mockFunc(&CreateBookDeps{
					repository: repository,
					trx:        trx,
					validator:  validator,
					cache:      cache,
				})
			}

			cb := usecase.NewCreateBook(repository, cache, validator)
			_, err := cb.Do(context.Background(), tt.createBookArgs)

			tt.wantErr(t, err)
		})
	}
}

func TestCreateBook_SadCase(t *testing.T) {
	testCases := []CreateBookTestCase{
		{
			name: "error author_id empty",
			createBookArgs: dto.Book{
				Title:       "some book",
				Description: "description",
				PublishDate: time.Now(),
			},
			mockFunc: func(md *CreateBookDeps) {
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(tt, err, utils.ErrInvalidAuthorId)
			},
		},
		{
			name: "validation error, expected return error",
			createBookArgs: dto.Book{
				Title:       "some book",
				Description: "description",
				PublishDate: time.Now(),
				Author: dto.Author{
					ID: 1,
				},
			},
			mockFunc: func(md *CreateBookDeps) {
				md.validator.EXPECT().Do(gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(tt, err, assert.AnError)
			},
		},
		{
			name: "duplicate title error, expected return error",
			createBookArgs: dto.Book{
				Title:       "some book",
				Description: "description",
				PublishDate: time.Now(),
				Author: dto.Author{
					ID: 1,
				},
			},
			mockFunc: func(md *CreateBookDeps) {
				md.validator.EXPECT().Do(gomock.Any(), gomock.Any()).Return(nil)
				md.repository.EXPECT().Begin(gomock.Any()).Return(md.trx, nil)
				md.repository.EXPECT().CreateBook(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(dto.Book{}, utils.ErrDuplicateBookTitle)
				md.trx.EXPECT().Rollback(gomock.Any()).Return(nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(tt, err, utils.ErrDuplicateBookTitle)
			},
		},
		{
			name: "rollback error, expected return error",
			createBookArgs: dto.Book{
				Title:       "some book",
				Description: "description",
				PublishDate: time.Now(),
				Author: dto.Author{
					ID: 1,
				},
			},
			mockFunc: func(md *CreateBookDeps) {
				md.validator.EXPECT().Do(gomock.Any(), gomock.Any()).Return(nil)
				md.repository.EXPECT().Begin(gomock.Any()).Return(md.trx, nil)
				md.repository.EXPECT().CreateBook(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(dto.Book{}, utils.ErrDuplicateBookTitle)
				md.trx.EXPECT().Rollback(gomock.Any()).Return(assert.AnError)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(tt, err, assert.AnError)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			repository := mocks.NewMockIRepository(c)
			trx := mocks.NewMockITransaction(c)
			validator := mocks.NewMockIValidator(c)
			cache := mocks.NewMockICache[dto.Book](c)

			if tt.mockFunc != nil {
				tt.mockFunc(&CreateBookDeps{
					repository: repository,
					trx:        trx,
					validator:  validator,
					cache:      cache,
				})
			}

			cb := usecase.NewCreateBook(repository, cache, validator)
			_, err := cb.Do(context.Background(), tt.createBookArgs)

			tt.wantErr(t, err)
		})
	}
}
