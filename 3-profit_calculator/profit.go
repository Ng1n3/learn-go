package main

import "fmt"

func main () {
	var taxRate float64
	var revenue float64
	var expenses float64

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("enter your tax rate: ")
	fmt.Scan(&taxRate)

	var EBT float64 = revenue - expenses
	var profit float64 = revenue - (revenue * (taxRate/100)) - expenses

	fmt.Println("My earnings before tax: ",EBT)
	fmt.Println("My earnings after tax: ", profit)
}