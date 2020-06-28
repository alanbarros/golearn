package main

import (
	"fmt"
	"log"

	"github.com/alanbarros/golearn/fullcycle/aula1/application/usecases"

	"github.com/alanbarros/golearn/fullcycle/aula1/application/repositories"
	"github.com/alanbarros/golearn/fullcycle/aula1/domain"
	_ "github.com/lib/pq"
)

func main() {

	user, err := domain.NewUser("Alan", "alan@alan1.com", "12345")

	if err != nil {
		log.Fatal(err)
	}

	useCase := usecases.UserUseCase{UserRepository: repositories.NewRepositoryDb()}

	err = useCase.Create(user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted:", user.Name, user.Email)
}
