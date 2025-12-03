package model

import (
	"time"

	"gorm.io/gorm"
)

// RefreshToken represents a persisted refresh token for a user.
type RefreshToken struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Token     string         `gorm:"type:varchar(512);not null;uniqueIndex" json:"token"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	ExpiresAt time.Time      `gorm:"not null;index" json:"expires_at"`
	Revoked   bool           `gorm:"default:false" json:"revoked"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// IsExpired returns true if the refresh token is expired.
func (r *RefreshToken) TableName() string {
	return "refresh_token"
}

func MigrateToken(db *gorm.DB) error {

	return db.AutoMigrate(&RefreshToken{})
}

func DropToken(db *gorm.DB) error {
	return db.Migrator().DropTable(&RefreshToken{})
}
