package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

  if err != nil {
    return 1000, errors.New("failed to find balance file")
  }
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

  if err != nil {
    return 1000, errors.New("failed to parse stored balance value")
  }
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

  if err != nil {
    fmt.Println("ERROR")
    fmt.Println(err)
    fmt.Println("----------")
    panic("Can't coninue, sorry.")
  }

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("your balance is", accountBalance)

		case 2:
			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. deposit must be greater than 0.")
				continue
			}

			accountBalance += depositAmount // accountbalance = accountBalance + depositAmount
			fmt.Println("Balance Updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Println("How much do you want to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. withdrawal must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than the balance")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New balance: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for reaching our Bank.")
			return
			// break
		}
	}
}
