package service
//not used but kept for reference
import (
	"context"

	"github.com/go-redis/redis/v8"
)

// BaseRepository defines minimal CRUD methods your repositories should have.
// This allows the BaseService to work with any repository type.
type BaseRepository[T any] interface {
	Create(ctx context.Context, model *T) error
	FindByID(ctx context.Context, id any) (*T, error)
	Update(ctx context.Context, model *T) error
	Delete(ctx context.Context, model *T) error
}

// BaseService provides reusable service logic for any model.
type BaseService[T any] struct {
	repo     BaseRepository[T]
	cache    *redis.Client
	useCache bool
}

// NewBaseService creates a new BaseService.
func NewBaseService[T any](repo BaseRepository[T], cache *redis.Client, useCache bool) *BaseService[T] {
	return &BaseService[T]{repo: repo, cache: cache, useCache: useCache}
}

// Create delegates to the repository.
func (s *BaseService[T]) Create(ctx context.Context, model *T) error {
	return s.repo.Create(ctx, model)
}

// FindByID delegates to the repository.
func (s *BaseService[T]) FindByID(ctx context.Context, id any) (*T, error) {
	return s.repo.FindByID(ctx, id)
}

// Update delegates to the repository.
func (s *BaseService[T]) Update(ctx context.Context, model *T) error {
	return s.repo.Update(ctx, model)
}

// Delete delegates to the repository.
func (s *BaseService[T]) Delete(ctx context.Context, model *T) error {
	return s.repo.Delete(ctx, model)
}
