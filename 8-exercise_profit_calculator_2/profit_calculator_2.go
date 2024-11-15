package main

import (
	"fmt"
	"os"
)

//  Goals
//  1) Validate user input
//     => Show error message & exit if invalid input is provided
//     - No negative numbers
//     - Not 0
//   2) Store calculated results into a file

const accountFile = "balance.txt"

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	
  // var accountBalance, err = getBalanceInfo()
	// if err != nil {
	// 	fmt.Println("----------")
	// 	fmt.Println("ERROR:", err)
	// 	fmt.Println("----------")
	// 	os.Exit(1)
	// }

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	writeFinancialsToFile(profit, ratio, ebt)
	return ebt, profit, ratio
}

func getUserInput(text string) float64 {
	var input float64
	for {
		fmt.Println(text)
		fmt.Scan(&input)
		if input > 0 {
			break
		}
		fmt.Println("Sorry only numbers above zero is required")

	}
	return input
}

func writeFinancialsToFile(profit, ratio, ebt float64) {
	balanctText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile(accountFile, []byte(balanctText), 0644)
}

// func getBalanceInfo() (float64, error) {
// 	data, err := os.ReadFile(accountFile)

// 	if err != nil {
// 		return 0, errors.New("failed to find account file")
// 	}

// 	balanceText := string(data)
// 	balance, err := strconv.ParseFloat(balanceText, 64)

// 	if err != nil {
// 		return 0, errors.New("failed to parse stored balance value")
// 	}

// 	return balance, nil
// }
