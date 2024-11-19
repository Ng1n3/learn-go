package main

import "fmt"

type Product struct {
  title string
  id string
  price float64
}

type TemperatureData struct {
  day1 float64
  day2 float64
}

func main() {
  var productNames [4]string = [4]string{"A book"}
  prices := [4]float64{2, 9.9, 6.9, 45.9}
  fmt.Println(prices)
  productNames[2] = "A Carpet"
  
  fmt.Println(productNames)
  fmt.Println(prices[2])
}