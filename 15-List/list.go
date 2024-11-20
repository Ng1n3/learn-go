package main

import "fmt"

// func main() {
//   var productNames [4]string = [4]string{"A book"}
//   prices := [4]float64{2, 9.9, 6.9, 45.9}
//   fmt.Println(prices)
//   productNames[2] = "A Carpet" 
//   fmt.Println(productNames)
//   fmt.Println(prices[2])
//   featuredPrices := prices[1:]
//   featuredPrices[0] = 80.0 
//   highlightedPrices := featuredPrices[:1]
//   fmt.Println(featuredPrices);
//   fmt.Println(highlightedPrices)
//   fmt.Println(prices)
//   fmt.Println(len(highlightedPrices), cap(highlightedPrices))
//   highlightedPrices = highlightedPrices[:3]
//   fmt.Println(highlightedPrices)
//   fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }

func main() {
  prices := []float64{10, 14.4}
  fmt.Println(prices[0:1])
  prices[1] = 9.5

  // updatedPrices := append(prices, 8.5, 12.9)
  // fmt.Println(updatedPrices, prices)

  prices = append(prices, 5.99)
  prices = prices[1:]
  fmt.Println(prices)
}