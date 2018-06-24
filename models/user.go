package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User - Model for declaring a user
type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `gorm:"type:varchar(100);unique"`
	FirstName string    `gorm:"size:255"`
	LastName  string    `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate ...
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

// BeforeUpdate - Updates Before
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
