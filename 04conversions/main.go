package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza App")
	fmt.Println("Rate our Pizza")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

	//conversion
	// used strings.Trimspace so that 1 can be added to the string, otherwise it would float error
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	//error handling
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating", numRating+1)
	}
}
