package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Color struct {
	ID          uuid.UUID `json:"id"`
	Events      []Event
	Name        string `gorm:"size:15"`
	Description string `gorm:"size:60"`
	Hex         string `gorm:"size:6;unique"`
}

// BeforeCreate ...
func (color *Color) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
