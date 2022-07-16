package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welc := "Welcome to user input"
	fmt.Println(welc)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter ur name: ")

	// comma ok || error ok syantac
	//input, err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

}
