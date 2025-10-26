package repository

import (
	"context"

	"gorm.io/gorm"
)

// BaseRepository provides basic CRUD operations for any model.
type BaseRepository[T any] struct {
	db *gorm.DB
}

// NewBaseRepository creates a new BaseRepository for any type T.
func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

// Create inserts a new record into the database.
func (r *BaseRepository[T]) Create(ctx context.Context, model *T) error {
	return r.db.WithContext(ctx).Create(model).Error
}

// FindByID retrieves a record by ID.
func (r *BaseRepository[T]) FindByID(ctx context.Context, id any) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &entity, nil
}

// Update updates a record.
func (r *BaseRepository[T]) Update(ctx context.Context, model *T) error {
	return r.db.WithContext(ctx).Save(model).Error
}

// Delete removes a record.
func (r *BaseRepository[T]) Delete(ctx context.Context, model *T) error {
	return r.db.WithContext(ctx).Delete(model).Error
}
