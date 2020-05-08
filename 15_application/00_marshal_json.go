package main

import "fmt"
import "encoding/json"

// This is the people type which will be "Marshaled"
type person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	// This is the first value to be "Marshaled"
	p1 := person{
		Firstname: "John",
		Lastname:  "Snow",
		Age:       22,
	}

	// This is the second value to be "Marshaled"
	p2 := person{
		Firstname: "Daenerys",
		Lastname:  "Targaryen",
		Age:       21,
	}

	// This is the person's array
	people := []person{p1, p2}

	// Printing the person's array
	fmt.Println("\nThey are the people in the program:")
	fmt.Println(people)

	// "Marshaling" the person's array
	people_bs, err := json.Marshal(people)
	// Verifying if any error when "Marshaling"
	if err != nil {
		fmt.Println(err)
	}

	// Printing the JSON generated after "Marshal", it is of type `[]byte` and must be converted to `string` to be printed correctly
	fmt.Println("\nThis is the people in JSON:")
	fmt.Println(string(people_bs))
}
