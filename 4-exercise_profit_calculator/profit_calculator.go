package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := profitCalculation(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n",ebt)
	fmt.Printf("%.1f\n",  profit)
	fmt.Printf("%.3f\n", ratio)

}

func getUserInput(message string) float64 {
	var input float64
	fmt.Print(message)
	fmt.Scan(&input)
	return input
}

func profitCalculation(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
