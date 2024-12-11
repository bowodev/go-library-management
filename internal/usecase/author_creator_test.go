package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/bowodev/go-library-management/internal/dto"
	mocks "github.com/bowodev/go-library-management/internal/interfaces/mock"
	"github.com/bowodev/go-library-management/internal/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateAuthor(t *testing.T) {
	testCases := []CreateAuthorTestCase{
		{
			name: "success",
			createAuthorArgs: dto.Author{
				Name:      "oe",
				Bio:       "oe oe oe",
				BirthDate: time.Now(),
			},
			mockFunc: func(md *CreateAuthorDeps) {
				md.validator.EXPECT().Do(gomock.Any(), gomock.Any()).Return(nil)
				md.repository.EXPECT().Begin(gomock.Any()).Return(md.trx, nil)
				md.repository.EXPECT().CreateAuthor(gomock.Any(), gomock.Any(), gomock.Any()).Return(dto.Author{}, nil)
				md.trx.EXPECT().Commit(gomock.Any()).Return(nil)
				md.cache.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
		{
			name: "success even redis failed to set cache",
			createAuthorArgs: dto.Author{
				Name:      "oe",
				Bio:       "oe oe oe",
				BirthDate: time.Now(),
			},
			mockFunc: func(md *CreateAuthorDeps) {
				md.validator.EXPECT().Do(gomock.Any(), gomock.Any()).Return(nil)
				md.repository.EXPECT().Begin(gomock.Any()).Return(md.trx, nil)
				md.repository.EXPECT().CreateAuthor(gomock.Any(), gomock.Any(), gomock.Any()).Return(dto.Author{}, nil)
				md.trx.EXPECT().Commit(gomock.Any()).Return(nil)
				md.cache.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
		{
			name: "failed to create author",
			createAuthorArgs: dto.Author{
				Name:      "oe",
				Bio:       "oe oe oe",
				BirthDate: time.Now(),
			},
			mockFunc: func(md *CreateAuthorDeps) {
				md.validator.EXPECT().Do(gomock.Any(), gomock.Any()).Return(nil)
				md.repository.EXPECT().Begin(gomock.Any()).Return(md.trx, nil)
				md.repository.EXPECT().CreateAuthor(gomock.Any(), gomock.Any(), gomock.Any()).Return(dto.Author{}, assert.AnError)
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
			cacher := mocks.NewMockICache[dto.Author](c)
			repository := mocks.NewMockIRepository(c)
			trx := mocks.NewMockITransaction(c)
			validator := mocks.NewMockIValidator(c)

			if tt.mockFunc != nil {
				tt.mockFunc(&CreateAuthorDeps{
					repository: repository,
					cache:      cacher,
					trx:        trx,
					validator:  validator,
				})
			}

			g := usecase.NewAuthorCreator(repository, cacher, validator)
			_, err := g.Do(context.Background(), tt.createAuthorArgs)

			tt.wantErr(t, err)
		})
	}
}
