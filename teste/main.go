package main

import (
	"github.com/alanbarros/golearn/teste/domain"
)

func main() {
	person := domain.NewPerson("Alan")

	person.Speak("Hello world!!")
}
