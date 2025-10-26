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
    db *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
    return &urlRepository{db: db}
}

func (r *urlRepository) Save(ctx context.Context, url *model.URL) error {
    return r.db.WithContext(ctx).Create(url).Error
}

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
