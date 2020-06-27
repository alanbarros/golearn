package main

import (
	"fmt"
	"log"

	"github.com/alanbarros/golearn/fullcycle/aula1/application/repositories"
	"github.com/alanbarros/golearn/fullcycle/aula1/domain"
	"github.com/alanbarros/golearn/fullcycle/aula1/infrastructure/database"
	_ "github.com/lib/pq"
)

func main() {

	db := database.ConnectDB()

	user := domain.User{
		Name:     "Alan",
		Email:    "alan@barros.com",
		Password: "123",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}
	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
