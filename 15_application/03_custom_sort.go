package main

import "fmt"
import "sort"

// Defining the type person
type person struct {
	firstname string
	lastname  string
	age       int
}

// Defining a type ByAge that will be used to implement the interface `Interface`, and it will be used to order a slice of person 'by age'
type ByAge []person

// The `Interface` method set, the type `ByAge` implements these methods to define the `sort.Sort` behavior
// `Len()` method defines the slice size
func (p ByAge) Len() int { return len(p) }

// `Less()` method defines a expression to determinate which of the two current elements in the slice is lesser than the other one
func (p ByAge) Less(i, j int) bool { return p[i].age < p[j].age }

// `Swap()` method swaps the two current elements
func (p ByAge) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	// Creating a slice of person with three elements
	people := []person{
		person{
			firstname: "Carmen",
			lastname:  "Herrera",
			age:       26,
		},
		person{
			firstname: "Alberto",
			lastname:  "Barros",
			age:       24,
		},
		person{
			firstname: "Emilio",
			lastname:  "Gutierrez",
			age:       25,
		},
	}

	// Printing the unordered slice of person
	fmt.Println("\nThe unordered slice of person:")
	fmt.Println(people)

	// Sorting the slice of person, it converts the slice of person to the type ByAge, then `Sort()` method use it to order the slice by age
	sort.Sort(ByAge(people))

	// Printing the slice of person ordered
	fmt.Println("\nThe ordered slice of person by age:")
	fmt.Println(people)
}
