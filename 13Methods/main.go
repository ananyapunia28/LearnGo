package main

import "fmt"

func main() {
	fmt.Println("Structs")

	ananya := user{"Ananya", "ananya@gmail.com", true, 20}
	fmt.Println(ananya)
	fmt.Println("Ananya's details: %+v\n", ananya)
	ananya.GetStatus()
	ananya.NewMail()

}

type user struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u user) GetStatus() {
	fmt.Println("is user active?", u.Status)
}

func (u user) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is", u.Email)
}
