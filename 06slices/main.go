package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Mango", "Peach"}
	fmt.Printf("Type of fruitList : %T\n", fruitList)

	fruitList = append(fruitList, "Banana", "Guava")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	//remove a value from slice based on index

	var courses = []string{"reactjs", "javascript", "swift", "python"}
	// fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
