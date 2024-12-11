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

func TestGetAuthor(t *testing.T) {
	testCases := []GetAuthorByIdTestCase{
		{
			name:              "when cache exists, get from redis",
			getAuthorByIdArgs: 1,
			mockFunc: func(md *GetAuthorByIdDeps) {
				md.cache.EXPECT().Get(gomock.Any(), gomock.Any()).Return(dto.Author{
					ID: 1,
				}, nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
		{
			name: "when redis error, get from db",
			mockFunc: func(md *GetAuthorByIdDeps) {
				md.cache.EXPECT().Get(gomock.Any(), gomock.Any()).Return(dto.Author{}, assert.AnError)
				md.repository.EXPECT().GetAuthorById(gomock.Any(), gomock.Any(), gomock.Any()).Return(dto.Author{
					ID: 1,
				}, nil)
				md.cache.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
		{
			name: "when cache empty, get from db",
			mockFunc: func(md *GetAuthorByIdDeps) {
				md.cache.EXPECT().Get(gomock.Any(), gomock.Any()).Return(dto.Author{}, nil)
				md.repository.EXPECT().GetAuthorById(gomock.Any(), gomock.Any(), gomock.Any()).Return(dto.Author{
					ID: 1,
				}, nil)
				md.cache.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			cacher := mocks.NewMockICache[dto.Author](c)
			repository := mocks.NewMockIRepository(c)

			if tt.mockFunc != nil {
				tt.mockFunc(&GetAuthorByIdDeps{
					repository: repository,
					cache:      cacher,
				})
			}

			g := usecase.NewAuthorGetter(repository, cacher)
			_, err := g.Do(context.Background(), tt.getAuthorByIdArgs)

			tt.wantErr(t, err)
		})
	}
}
