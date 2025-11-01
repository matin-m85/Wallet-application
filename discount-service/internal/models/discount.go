package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Discount struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Code        string    `gorm:"size:64;uniqueIndex;not null" json:"code"`
	Description string    `gorm:"size:255" json:"description"`
	Percent     int       `gorm:"not null" json:"percent"`
	Remaining   int       `gorm:"not null;default:1" json:"remaining"`
	IsActive    bool      `gorm:"not null;default:true" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (d *Discount) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return nil
}
