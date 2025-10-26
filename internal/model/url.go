package model

import "time"

type URL struct {
    ID        uint      `gorm:"primaryKey"`
    ShortCode string    `gorm:"uniqueIndex;not null"`
    Original  string    `gorm:"not null"`
    CreatedAt time.Time
}
