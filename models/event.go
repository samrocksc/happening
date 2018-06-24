package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Event - Model for an event
type Event struct {
	ID          uuid.UUID `json:"id"`
	ColorID     int       `gorm:"index"`
	Name        string    `gorm:"size:50"`
	Description string    `gorm:"size:500"`
	Lat         int32
	Long        int32
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

// BeforeCreate ...
func (event *Event) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

// BeforeUpdate - Updates Before
func (event *Event) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
