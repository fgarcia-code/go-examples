package main

import "fmt"

// Declare struct of type person
type person struct {
	 firstname string
	 lastname string
}

type professional struct {
	person
	profession string
}

func main() {
	// Declare variable p1 of type person
	p1 := person{
		firstname: "Juan",
		lastname: "Perez",
	}

	// Declare variable p2 of type person
	p2 := person{
		firstname: "Jose",
		lastname: "Sandoval",
	}

	// Accessing the elements of the struct, and print it
	fmt.Println("\nThe people are:")
	fmt.Println(p1.firstname, p1.lastname)
	fmt.Println(p2.firstname, p2.lastname)

   // Embedded struct
   p3 := professional{
   	person: p2,
   	profession: "Data Engineer",
   }
   fmt.Println("\nThe professionals are:")
   fmt.Println(p3.firstname,p3.lastname,p3.profession)
}