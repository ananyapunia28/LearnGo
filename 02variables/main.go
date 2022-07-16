package main

import "fmt"

//first letter capital = public keyword
const LoginToken string = "jhbdfhsd"

func main() {
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45453253456345
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//implicit type

	var website = "dsckiet.com"
	fmt.Println(website)

	//no var style
	//important one
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
