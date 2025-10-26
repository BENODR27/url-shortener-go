package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/BENODR27/url-shortener-go/internal/config"
	"github.com/BENODR27/url-shortener-go/internal/handler"
	"github.com/BENODR27/url-shortener-go/internal/model"
	"github.com/BENODR27/url-shortener-go/internal/repository"
	"github.com/BENODR27/url-shortener-go/internal/service"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize DB
	var db *gorm.DB
	var err error
	if cfg.DBDriver == "sqlite" {
		db, err = gorm.Open(sqlite.Open(cfg.DSN()), &gorm.Config{})
	} else {
		db, err = gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	}
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Auto-migrate URL table
	if err := db.AutoMigrate(&model.URL{}); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	// Initialize Redis (optional)
	var cache *redis.Client
	if cfg.UseRedis {
		cache = redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
		if err := cache.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Sprintf("failed to connect to Redis: %v", err))
		}
		fmt.Println("Redis caching enabled")
	} else {
		fmt.Println("Redis caching disabled")
	}

	// Initialize repository and service
	repo := repository.NewURLRepository(db)
	svc := service.NewURLService(repo, cache, cfg.UseRedis)

	// Initialize handler
	h := handler.NewURLHandler(svc)

	// Setup Gin router
	r := gin.Default()
	r.POST("/shorten", h.Shorten)
	r.GET("/:code", h.Resolve)

	// Run server
	fmt.Printf("Server running on port %s\n", cfg.ServerPort)
	if err := r.Run(fmt.Sprintf(":%s", cfg.ServerPort)); err != nil {
		panic(err)
	}
}
