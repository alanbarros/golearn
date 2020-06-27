package domain_test

import (
	"testing"

	"github.com/alanbarros/golearn/fullcycle/aula1/domain"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {

	_, err := domain.NewUser("Alan", "alan@alan.com", "1234")
	require.NoError(t, err)

	_, err = domain.NewUser("Alan", "alanalan.com", "1234")
	require.Error(t, err)

	_, err = domain.NewUser("Alan", "alan@alan.com", "")
	require.Error(t, err)

	_, err = domain.NewUser("Alan", "alan@alan.com", "")
	require.Error(t, err)
}
