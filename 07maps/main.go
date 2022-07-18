package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List Of All Languages: ", languages)
	fmt.Println("JS: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Langauges: ", languages)

}
