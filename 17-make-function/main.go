package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
  fmt.Println(m)
}

func main() {
  userNames := make([]string, 2, 5)

  userNames[0] = "Julie"
  userNames = append(userNames, "max")
  userNames = append(userNames, "Manuel")

  fmt.Println(userNames)

  courseRating := make(floatMap, 3)

  courseRating["go"] = 4.8
  courseRating["react"] = 4.5
  courseRating["angular"] = 4.7
  courseRating["node"] = 4.9

  courseRating.output()
  // fmt.Println(courseRating)

  for index, value := range userNames {
    fmt.Println("Index:", index)
    fmt.Println("Value:", value)
  }

  for key, value := range courseRating {
    fmt.Println("Key:", key)
    fmt.Println("Value:", value)
  }
}