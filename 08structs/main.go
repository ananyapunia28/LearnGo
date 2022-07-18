package main

import "fmt"

func main() {
	fmt.Println("Structs")

	ananya := user{"Ananya", "ananya@gmail.com", true, 20}
	fmt.Println(ananya)
	fmt.Println("Ananya's details: %+v\n", ananya)

}

type user struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
