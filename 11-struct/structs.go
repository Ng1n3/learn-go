package main

import (
	"fmt"
  "example.com/struct/user"
)



// func outputUserDetails(user *User) {
//   fmt.Println((*user).firstName, user.lastName, user.birthDate)
// }


func main() {
  userFirstName := getUserData("please enter your first name: ")
  userLastName := getUserData("Please enter your last name: ")
  userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYY): ")

  var appUser *user.User

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

  // appUser, err  := newuser(userFirstName, userLastName, userBirthdate)
  appUser, err := user.New(userFirstName, userLastName, userBirthdate)
  if err != nil {
    fmt.Println(err)
    return
  }

  admin := user.NewAdmin("test@example.com", "test123")

  admin.OutputUserDetails()
  admin.ClearUserName()

  appUser.OutputUserDetails()
  appUser.ClearUserName()
  appUser.OutputUserDetails()
}


func getUserData(prompText string) string {
  fmt.Print(prompText)
  var value string
  fmt.Scanln(&value)
  return value
}