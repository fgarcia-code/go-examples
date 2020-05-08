package main

import "fmt"

// The interface defines a type of multiple types, to be valid they must implement write method
type writer interface {
	write()
}

// Programmer implements the write method
type programmer struct {
	firstname string
	lastname string
}

func (p programmer) write()  {
	fmt.Println(p.firstname,p.lastname,"says:")
	fmt.Println("I'm writing code\n")
}

// Poet implements the write method
type poet struct {
	firstname string
	lastname string
}

func (p poet) write()  {
	fmt.Println(p.firstname,p.lastname,"says:")
	fmt.Println("I'm writing a poem\n")
}

// The function typing use the interface writer and invokes the method write
func typing(w writer)  {
	switch w.(type) {
	case poet:
		fmt.Println("The poet",w.(poet).firstname)
	case programmer:
		fmt.Println("The programmer",w.(programmer).firstname)
	}
	w.write()
}

func main() {
	// We create a variable of type programer
	p1 := programmer{
		firstname: "Linus",
		lastname: "Torvalds",
	}

	// We create a variable of type poet
	p2 := poet{
		firstname: "Pablo",
		lastname: "Neruda",
	}

	// The method can receive both types, `poet` and `programmer` because they are writers
	typing(p1)
	typing(p2)
}