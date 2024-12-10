package interfaces

import "context"

//go:generate mockgen -source=cache.go -destination=./mock/cache_mock.go -package=mocks
type ICache[T any] interface {
	Set(ctx context.Context, key string, value T) error
	Get(ctx context.Context, key string) (T, error)
	Del(ctx context.Context, key string) error
}
