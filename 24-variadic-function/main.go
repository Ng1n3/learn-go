package main

import "fmt"

func main() {
  // numbers := []int{1, 10, 15}
  sum := sumup(1, 10, 15, 28 -50)

  fmt.Println(sum)
}

func sumup(numbers ...int) int {
  sum := 0

  for _, val := range numbers {
    sum += val
  }

  return sum
}