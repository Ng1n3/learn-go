package main

import (
	"fmt"
	"time"
)

type User struct {
  firstName string
  lastName string
  birthDate string
  createdAt time.Time
}

func main() {
  userFirstName := getUserData("please enter your first name: ")
  userLastName := getUserData("Please enter your last name: ")
  userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYY): ")

  var appUser User

  // appUser = User{
  //   firstName: userFirstName,
  //   lastName: userLastName,
  //   birthDate: userBirthdate,
  //   createdAt: time.Now(),
  // }
  // appUser = User{}

  appUser = User{
    userFirstName,
    userLastName,
    userBirthdate,
    time.Now(),
  }

  // ... do something awesome with that generated data!

  outputUserDetails(appUser)
}

func outputUserDetails(user User) {
  fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func getUserData(prompText string) string {
  fmt.Print(prompText)
  var value string
  fmt.Scan(&value)
  return value
}