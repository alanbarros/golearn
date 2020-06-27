package domain

import "fmt"

// Person is an entity of domain
type Person struct {
	Base
	Name string
}

// NewPerson make a new instance of Person
func NewPerson(name string) Person {
	return Person{
		Base: newBase(),
		Name: name,
	}
}

// Speak make a peaple say something
func (person Person) Speak(anything string) {

	if anything == "" {
		fmt.Printf("The person %v says nothing \n", person.Name)
	}

	fmt.Printf("The person %v says: %v \n", person.Name, anything)
}
