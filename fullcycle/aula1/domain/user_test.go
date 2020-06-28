package domain_test

import (
	"testing"

	"github.com/alanbarros/golearn/fullcycle/aula1/domain"
	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	_, err := domain.NewUser("Wesley", "wesleyxwesleyxxxgmail.com", "12345678")
	require.Error(t, err)

	_, err = domain.NewUser("Wesley", "wesleyxwesleyxxxgmail.com", "")
	require.Error(t, err)

	_, err = domain.NewUser("", "wesleyxwesleyxxxgmail.com", "")
	require.Error(t, err)

	user, err := domain.NewUser("Wesley", "wesleyxwesleyxxx@gmail.com", "12345678")
	require.Nil(t, err)

	govalidator.IsUUIDv4(user.Token)

}
