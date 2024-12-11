package usecase_test

import (
	"context"
	"testing"

	"github.com/bowodev/go-library-management/internal/dto"
	mocks "github.com/bowodev/go-library-management/internal/interfaces/mock"
	"github.com/bowodev/go-library-management/internal/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteAuthor(t *testing.T) {
	testCases := []DeleteAuthorTestCase{
		{
			name:             "success",
			deleteAuthorArgs: 1,
			mockFunc: func(md *DeleteAuthorDeps) {
				md.repository.EXPECT().Begin(gomock.Any()).Return(md.trx, nil)
				md.repository.EXPECT().GetBookByAuthor(gomock.Any(), gomock.Any(), gomock.Any()).Return(dto.Books{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
				}, nil)
				md.repository.EXPECT().DeleteAuthor(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				md.trx.EXPECT().Commit(gomock.Any()).Return(nil)
				md.authorCache.EXPECT().Del(gomock.Any(), gomock.Any()).Return(nil)
				md.bookCache.EXPECT().Del(gomock.Any(), gomock.Any()).Return(nil).Times(2)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
		{
			name:             "success even redis failed to set cache",
			deleteAuthorArgs: 1,
			mockFunc: func(md *DeleteAuthorDeps) {
				md.repository.EXPECT().Begin(gomock.Any()).Return(md.trx, nil)
				md.repository.EXPECT().GetBookByAuthor(gomock.Any(), gomock.Any(), gomock.Any()).Return(dto.Books{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
				}, nil)
				md.repository.EXPECT().DeleteAuthor(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				md.trx.EXPECT().Commit(gomock.Any()).Return(nil)
				md.authorCache.EXPECT().Del(gomock.Any(), gomock.Any()).Return(assert.AnError)
				md.bookCache.EXPECT().Del(gomock.Any(), gomock.Any()).Return(nil).Times(2)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
		{
			name:             "failed delete",
			deleteAuthorArgs: 1,
			mockFunc: func(md *DeleteAuthorDeps) {
				md.repository.EXPECT().Begin(gomock.Any()).Return(md.trx, nil)
				md.repository.EXPECT().GetBookByAuthor(gomock.Any(), gomock.Any(), gomock.Any()).Return(dto.Books{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
				}, nil)
				md.repository.EXPECT().DeleteAuthor(gomock.Any(), gomock.Any(), gomock.Any()).Return(assert.AnError)
				md.trx.EXPECT().Rollback(gomock.Any()).Return(nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(tt, err, assert.AnError)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			authorCacher := mocks.NewMockICache[dto.Author](c)
			bookCacher := mocks.NewMockICache[dto.Book](c)
			repository := mocks.NewMockIRepository(c)
			trx := mocks.NewMockITransaction(c)

			if tt.mockFunc != nil {
				tt.mockFunc(&DeleteAuthorDeps{
					repository:  repository,
					trx:         trx,
					authorCache: authorCacher,
					bookCache:   bookCacher,
				})
			}

			g := usecase.NewAuthorDeleter(repository, authorCacher, bookCacher)
			err := g.Do(context.Background(), tt.deleteAuthorArgs)

			tt.wantErr(t, err)
		})
	}
}
