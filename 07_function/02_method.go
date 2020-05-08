package main

import "fmt"

type person struct {
	firstname string
	lastname string
}

type professional struct {
	person
	profession string
}

// func (r recevier) indentifier(parameters) (return(s)) { ... }
func (s professional) speak()  {
	// Accessing the field professional.profession
	fmt.Println("Hello, I'm a", s.profession)
	fmt.Println()
}

func main() {
	p1 := professional{
		person: person{
			"Fernando",
			"Gomez",

		},
		profession: "Data Engineer",
	}

	// Calling the method speak from professional struct
	fmt.Println("\nCalling the method speak:")
	p1.speak()
}