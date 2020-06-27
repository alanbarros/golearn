package database

import (
	"log"
	"os"

	"github.com/alanbarros/golearn/fullcycle/aula1/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// ConnectDB make a Db connection
func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loding .env file")
	}

	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	//defer db.Close()

	db.AutoMigrate(&domain.User{})

	return db
}
