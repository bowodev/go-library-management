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

func TestGetAllAuthor(t *testing.T) {
	testCases := []GetAuthorAllTestCase{
		{
			name: "success get from db",
			mockFunc: func(md *GetAuthorAllDeps) {
				md.repository.EXPECT().GetAuthorAll(gomock.Any(), gomock.Any()).Return(dto.Authors{}, nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			cacher := mocks.NewMockICache[dto.Authors](c)
			repository := mocks.NewMockIRepository(c)

			if tt.mockFunc != nil {
				tt.mockFunc(&GetAuthorAllDeps{
					repository: repository,
					cache:      cacher,
				})
			}

			g := usecase.NewAuthorAllGetter(repository)
			_, err := g.Do(context.Background())

			tt.wantErr(t, err)
		})
	}
}
