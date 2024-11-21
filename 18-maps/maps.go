package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon web Services"])

	websites["Udemy"] = "https://udemy.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}

/*
[Lecture]:
  - You cannot delete a key value pair when using structs unlike when using maps
[lecture]:
  - You cannot add to a key value pair when using structs, everything has to be predefined from the start. Those are the differences between a struct and a map
*/
