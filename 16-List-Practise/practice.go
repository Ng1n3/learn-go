package main

import "fmt"

type Product struct {
	Id    string
	Title string
	Price float64
}

func main() {
	hobbies := [3]string{"Coding", "Games", "Tech"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	newHobbies := hobbies[:2]
	newHobbies2 := hobbies[0:2]
	fmt.Println(newHobbies)
	fmt.Println(newHobbies2)

	reSliced := [2]string{hobbies[1], hobbies[2]}
	fmt.Println(reSliced)

	goals := []string{"Become a professional Go Developer", "Improve my backend skills"}
	goals[1] = "An Improved backend Engineer"
	goals = append(goals, "Improve my knowledge of programming")
	fmt.Println(goals)

	newList := []Product{{Id: "adfads124", Title: "Effective Go", Price: 60.0}, {Id: "dafda23r3", Title: "Mastering Go", Price: 45.0}}
	fmt.Println(newList)

	newList = append(newList, Product{Id: "dafdssafd", Title: "Welcome to Go", Price: 70.0})

  fmt.Println(newList)
}

// Time to practice what you learned!

/*
1) Create a new array (!) that contains three hobbies you have Output (print) that array in the command line.
2) Also output more data about that array:
  - The first element (standalone)
  - The second and third element combined a new list
3) Creat a slice based on the first element that contains the first and second elements.
    Create that slice in two different ways (i.e. create two slices in the end)
4) Re-slice the slice from (3) and change it to contain the second and last element of the original array.
5) Create a "dynamic array" that contains your course goals (at least 2 goals)
6) Set the second goal to a different one AND then add a third goal to that existing
7) Bonus: Create a "Product" struct with title, id, price and create a dynamic list of products (at least 2 products).
  Then add a third product to the existing list of products.

*/
