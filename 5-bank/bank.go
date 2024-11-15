package main

import "fmt"



func main() {
  var accountBalance = 1000.0

  for i := 0; i < 2; i++ {
    fmt.Println("Welcome to Go Bank!")
    fmt.Println("What do you want to do?")
    fmt.Println("1. Check balance")
    fmt.Println("2. Deposit money")
    fmt.Println("3. Withdraw money")
    fmt.Println("4. Exit")
  
    var choice int
    fmt.Print("Your choice: ")
    fmt.Scan(&choice)
  
    // wantsCheckBalance := choice == 1
  
    if choice == 1 {
       fmt.Println("your balance is", accountBalance)
    } else if choice == 2 {
      fmt.Println("Your deposit: ")
      var depositAmount float64
      fmt.Scan(&depositAmount)
  
      if depositAmount <= 0 {
        fmt.Println("Invalid amount. deposit must be greater than 0.")
        return
      }
  
      accountBalance += depositAmount // accountbalance = accountBalance + depositAmount
      fmt.Println("Balance Updated! New amount: ", accountBalance);
    } else if choice == 3 {
      fmt.Println("How much do you want to withdraw: ")
      var withdrawAmount float64
      fmt.Scan(&withdrawAmount)
      if withdrawAmount <= 0 {
        fmt.Println("Invalid amount. withdrawal must be greater than 0.")
      }
  
      if withdrawAmount > accountBalance {
        fmt.Println("Invalid amount. You can't withdraw more than the balance")
      }
  
      accountBalance -= withdrawAmount
      fmt.Println("Balance updated! New balance: ", accountBalance)
    } else {
      fmt.Println("Goodbye!")
    }

  }

  // fmt.Println("Your choice: ", choice)
}