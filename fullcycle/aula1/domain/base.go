package domain

import (
	"time"
)

// Base represents a default entity
type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key" valid:"notnull,uuid"`
	CreatedAt time.Time `json:"created_at" valid:"required"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}
