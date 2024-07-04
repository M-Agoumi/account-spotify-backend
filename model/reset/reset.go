package reset

import (
	"gorm.io/gorm"
	"time"
)

type PasswordReset struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"size:255;not null;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	ExpiresAt time.Time `gorm:"not null"`
}
