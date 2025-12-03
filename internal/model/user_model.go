package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents an application user.
type User struct {
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID         string         `gorm:"type:char(36);uniqueIndex;not null" json:"uuid"`
	Name         string         `gorm:"size:255;not null" json:"name"`
	Email        string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	PasswordHash string         `gorm:"size:255;not null" json:"-"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	RefreshToken []RefreshToken
	Link         []Link
}

// TableName returns the database table name for User.
func (User) TableName() string {
	return "users"
}

// BeforeCreate is a GORM hook that ensures a UUID is set before inserting.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UUID == "" {
		u.UUID = uuid.New().String()
	}
	return nil
}

// MigrateUsers runs the auto-migration for the User model.
func MigrateUsers(db *gorm.DB) error {

	return db.AutoMigrate(&User{})
}

func DropUsers(db *gorm.DB) error {
	return db.Migrator().DropTable(&User{})
}
