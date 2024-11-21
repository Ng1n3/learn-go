package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
  result := Add(2.4, 5)
  fmt.Println(result)

  result2 := Multiply(8, 8.3)
  fmt.Println(result2)

}

func Add[T constraints.Ordered](a, b T) T {
  return a + b
}

type Num interface {
  int | int32 | int16 | int8 | int64 | float32 | float64
}

func Multiply[T Num](a, b T) T {
  return a * b
}