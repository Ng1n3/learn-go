package main

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


func (user User) outputUserDetails() {
  
  fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func (user *User) clearUserName() {
  user.firstName = ""
  user.lastName = ""
}

func newuser(firstName, lastName, birthdate string) (*User, error) {
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

// func outputUserDetails(user *User) {
//   fmt.Println((*user).firstName, user.lastName, user.birthDate)
// }


func main() {
  userFirstName := getUserData("please enter your first name: ")
  userLastName := getUserData("Please enter your last name: ")
  userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYY): ")

  var appUser *User

  // appUser = User{
  //   firstName: userFirstName,
  //   lastName: userLastName,
  //   birthDate: userBirthdate,
  //   createdAt: time.Now(),
  // }
  // appUser = User{}
  // appUser = User{
  //   userFirstName,
  //   userLastName,
  //   userBirthdate,
  //   time.Now(),
  // }
  // ... do something awesome with that generated data!
  // outputUserDetails(&appUser)

  appUser, err  := newuser(userFirstName, userLastName, userBirthdate)

  if err != nil {
    fmt.Println(err)
    return
  }

  appUser.outputUserDetails()
  appUser.clearUserName()
  appUser.outputUserDetails()
}


func getUserData(prompText string) string {
  fmt.Print(prompText)
  var value string
  fmt.Scanln(&value)
  return value
}