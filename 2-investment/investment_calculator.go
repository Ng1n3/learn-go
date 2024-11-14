package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64 = 1000
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// fmt.Print("expected Return Rate: ")
	outputText("expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount *math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years) 
	
	// fmt.Println("Your future Value: ",futureValue)
	// fmt.Println("You real Value with inflation adjusted for: ", futureRealValue)
	// fmt.Printf("Future Value: %v\nFuture Value(adjusted for Inflation): %v", futureValue, futureRealValue)
	// fmt.Printf("Future Value: %f\nFuture Value(adjusted for Inflation): %f", futureValue, futureRealValue)
	// fmt.Printf("Future Value: %.0f\nFuture Value(adjusted for Inflation): %.0f", futureValue, futureRealValue)
	// fmt.Printf(`Future Value: %.0f
	// Future Value(adjusted for Inflation): %.0f`, futureValue, futureRealValue)
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formatedRFV := fmt.Sprintf("Future value (adjusted for Inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV,formatedRFV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount *math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return fv, rfv
	// return
}