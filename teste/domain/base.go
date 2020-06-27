package domain

import (
	uuid "github.com/satori/go.uuid"
)

// Base contains commons properties
type Base struct {
	ID string
}

func newBase() Base {
	guid, _ := uuid.NewV4()

	return Base{
		ID: guid.String(),
	}
}
