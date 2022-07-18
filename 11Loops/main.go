package main

import "fmt"

func main() {
	fmt.Println("Loops")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Println("Index is %v and value is %v", index, day)
	}

	rougueValue := 1
	for rougueValue < 10 {

		// if rougueValue == 5 {
		// 	break
		// }
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		if rougueValue == 2 {
			goto lco
		}

		fmt.Println("Value is :", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jump at LCO")
}
