package usecases

import (
	"log"

	"github.com/alanbarros/golearn/fullcycle/aula1/application/repositories"
	"github.com/alanbarros/golearn/fullcycle/aula1/domain"
)

// UserUseCase manage user flow
type UserUseCase struct {
	UserRepository repositories.UserRepository
}

// Create a new user
func (usecase *UserUseCase) Create(user *domain.User) error {

	_, err := usecase.UserRepository.Insert(user)

	if err != nil {
		log.Fatalf("Error to create user %v:", user)
		return err
	}

	return nil
}
