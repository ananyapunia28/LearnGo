package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Length of Fruit List: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg List is : ", len(vegList))
}
