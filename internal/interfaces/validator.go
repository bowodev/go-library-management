package interfaces

import "context"

//go:generate mockgen -source=validator.go -destination=./mock/validator_mock.go -package=mocks
type IValidator interface {
	Do(ctx context.Context, in interface{}) error
}
