package domain

import (
	"fmt"
	"log"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User is an domain entity
type User struct {
	Base     `valid:"required"`
	Name     string `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Email    string `json:"email" gorm:"type:varchar(255); unique_index" valid:"notnull,email"`
	Password string `json:"-" gorm:"type:varchar(255)" valid:"notnull"`
	Token    string `json:"token" gorm:"type:varchar(255); unique_index" valid:"notnull,uuid"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// NewUser provides a new instance of User
func NewUser(name string, email string, password string) (*User, error) {

	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.Prepare()

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Prepare is used to generate a encypted password and make a token
func (user *User) Prepare() error {

	if user.Password == "" {
		return fmt.Errorf("Password cannot be empty")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Error during the password generation: %v", err)
		return err
	}

	token, _ := uuid.NewV4()
	id, _ := uuid.NewV4()

	user.ID = id.String()
	user.CreatedAt = time.Now()
	user.Token = token.String()
	user.Password = string(password)

	err = user.validate()

	if err != nil {
		log.Fatalf("Error diring the user validation: %v", err)
		return err
	}

	return nil

}

func (user *User) validate() error {

	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}
