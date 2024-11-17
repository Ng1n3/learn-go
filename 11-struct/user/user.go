package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
  firstName string
  lastName string
  birthDate string
  createdAt time.Time
}

type Admin struct {
  email string
  password string
  User
}


func (user User) OutputUserDetails() {
  
  fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func (user *User) ClearUserName() {
  user.firstName = ""
  user.lastName = ""
}

func NewAdmin(email, password string) Admin {
  return Admin {
    email: email,
    password: password,
    User: User {
      firstName: "ADMIN",
      lastName: "ADMIN",
      birthDate: "----",
      createdAt: time.Now(),
    },
  }
}

func New(firstName, lastName, birthdate string) (*User, error) {
  if firstName == "" || lastName == "" || birthdate == "" {
    return nil, errors.New("first name, last name, birthdate are required")
  }


  return &User{
    firstName: firstName,
    lastName: lastName,
    birthDate: birthdate,
    createdAt: time.Now(),
  }, nil
}