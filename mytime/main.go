package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to golang time study")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))

}
