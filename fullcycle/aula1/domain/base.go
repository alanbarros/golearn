package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Base represents a default entity
type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" sql:"index"`
}

// BeforeCreate is used to create default colluns on dataBase
func (base *Base) BeforeCreate(scope *gorm.Scope) error {

	err := scope.SetColumn("CreatedAt", time.Now())

	guid, _ := uuid.NewV4()

	err = scope.SetColumn("ID", guid.String())

	if err != nil {
		log.Fatalf("Error during obj creating: %v", err)
		return err
	}

	return nil
}
