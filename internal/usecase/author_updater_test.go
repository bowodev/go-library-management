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

func TestUpdateAuthor(t *testing.T) {
	testCases := []UpdateAuthorTestCase{
		{
			name: "success",
			updateAuthorArgs: dto.Author{
				ID:        1,
				Name:      "oe",
				Bio:       "oe oe oe",
				BirthDate: time.Now(),
			},
			mockFunc: func(md *UpdateAuthorDeps) {
				md.validator.EXPECT().Do(gomock.Any(), gomock.Any()).Return(nil)
				md.repository.EXPECT().Begin(gomock.Any()).Return(md.trx, nil)
				md.repository.EXPECT().UpdateAuthor(gomock.Any(), gomock.Any(), gomock.Any()).Return(dto.Author{}, nil)
				md.trx.EXPECT().Commit(gomock.Any()).Return(nil)
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
			trx := mocks.NewMockITransaction(c)
			validator := mocks.NewMockIValidator(c)

			if tt.mockFunc != nil {
				tt.mockFunc(&UpdateAuthorDeps{
					repository: repository,
					cache:      cacher,
					trx:        trx,
					validator:  validator,
				})
			}

			g := usecase.NewAuthorUpdater(repository, cacher, validator)
			_, err := g.Do(context.Background(), tt.updateAuthorArgs)

			tt.wantErr(t, err)
		})
	}
}
