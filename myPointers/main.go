package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value", ptr)
	fmt.Println("Value of ptr = ", *ptr)

	*ptr = *ptr * 5
	fmt.Println("New value :", myNumber)
}
