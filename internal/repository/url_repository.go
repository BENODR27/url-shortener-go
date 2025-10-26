package repository

import (
	"context"
	"errors"

	"github.com/BENODR27/url-shortener-go/internal/model"
	"gorm.io/gorm"
)

type URLRepository interface {
	Save(ctx context.Context, url *model.URL) error
	FindByCode(ctx context.Context, code string) (*model.URL, error)
}

type urlRepository struct {
	*BaseRepository[model.URL] // Embed base repo
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &urlRepository{
		BaseRepository: NewBaseRepository[model.URL](db),
		db:             db,
	}
}

// Save overrides Create (can include business logic if needed)
func (r *urlRepository) Save(ctx context.Context, url *model.URL) error {
	return r.Create(ctx, url)
}

// FindByCode is URL-specific logic
func (r *urlRepository) FindByCode(ctx context.Context, code string) (*model.URL, error) {
	var url model.URL
	if err := r.db.WithContext(ctx).Where("short_code = ?", code).First(&url).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &url, nil
}
