package main

import "fmt"

func main() {
	fmt.Println("IfElse")

	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10 Login Count"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Not less than 10")
	}

}
