package govalidator

import (
	"context"

	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/go-playground/validator/v10"
)

type goValidator struct {
	validator *validator.Validate
}

var _ interfaces.IValidator = (*goValidator)(nil)

func New() *goValidator {
	v := validator.New()
	return &goValidator{
		validator: v,
	}
}

// Do implements interfaces.IValidator.
func (g *goValidator) Do(ctx context.Context, in interface{}) error {
	return g.validator.StructCtx(ctx, in)
}
