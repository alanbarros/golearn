package domain_test

import (
	"testing"

	"github.com/alanbarros/golearn/fullcycle/aula1/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	_, err := domain.NewUser("Alan", "alan@alan.com", "12345")

	assert.Nil(t, err)

}
