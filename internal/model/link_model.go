package model

import (
	"time"

	"gorm.io/gorm"
)

// Link represents a shortlink record for the shortlink application.
type Link struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Slug        string         `gorm:"size:100;uniqueIndex;not null" json:"slug"`
	UserID      uint           `gorm:"not null;index" json:"user_id"`
	Destination string         `gorm:"type:text;not null" json:"destination"`
	Title       string         `gorm:"size:255" json:"title"`
	Visits      uint64         `gorm:"default:0;not null" json:"visits"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName overrides the default table name.
func (Link) TableName() string {
	return "links"
}

// Migrate automigrates the Link model schema.
func MigrateLink(db *gorm.DB) error {
	return db.AutoMigrate(&Link{})
}

func DropLink(db *gorm.DB) error {
	return db.Migrator().DropTable(&Link{})
}
