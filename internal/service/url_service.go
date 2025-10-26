package service

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/BENODR27/url-shortener-go/internal/model"
	"github.com/BENODR27/url-shortener-go/internal/repository"
	"github.com/BENODR27/url-shortener-go/pkg/shortener"
)

// URLService defines the interface for URL operations
type URLService interface {
	Shorten(ctx context.Context, original string) (string, error)
	Resolve(ctx context.Context, code string) (string, error)
}

// urlService implements URLService
type urlService struct {
	repo     repository.URLRepository
	cache    *redis.Client
	useCache bool
}

// NewURLService creates a new URLService instance
func NewURLService(repo repository.URLRepository, cache *redis.Client, useCache bool) URLService {
	return &urlService{repo: repo, cache: cache, useCache: useCache}
}

// Shorten generates a short code and saves the URL
func (s *urlService) Shorten(ctx context.Context, original string) (string, error) {
	code := shortener.Generate(8)
	url := &model.URL{ShortCode: code, Original: original}
	if err := s.repo.Save(ctx, url); err != nil {
		return "", err
	}

	if s.useCache && s.cache != nil {
		_ = s.cache.Set(ctx, code, original, 0).Err()
	}

	return code, nil
}

// Resolve retrieves the original URL from the code
func (s *urlService) Resolve(ctx context.Context, code string) (string, error) {
	if s.useCache && s.cache != nil {
		if val, err := s.cache.Get(ctx, code).Result(); err == nil {
			return val, nil
		}
	}

	url, err := s.repo.FindByCode(ctx, code)
	if err != nil || url == nil {
		return "", errors.New("not found")
	}

	if s.useCache && s.cache != nil {
		_ = s.cache.Set(ctx, code, url.Original, 0).Err()
	}

	return url.Original, nil
}
