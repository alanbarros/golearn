package repositories

import (
	"log"

	"github.com/alanbarros/golearn/fullcycle/aula1/domain"
	"github.com/jinzhu/gorm"
)

// UserRepository is an interface to manipulate database
type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

// UserRepositoryDb implements UserRepository
type UserRepositoryDb struct {
	Db *gorm.DB
}

// Insert try Create a user into database
func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {
	err := user.Prepare()

	if err != nil {
		log.Fatalf("Error during user validation: %v", err)
		return user, err
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Error to persist user: %v", err)
		return user, err
	}

	return user, nil
}
