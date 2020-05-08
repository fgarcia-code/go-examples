package main

import "fmt"


func main() {
	p1 := struct {
		name string
		city string
	}{
		name: "Bellas Artes",
		city: "Mexico",
	}

	// This is a anonymous struct
	fmt.Println("\nThis is a anonymous Struct:")
	fmt.Println(p1)
	fmt.Printf("%T\n",p1)
}
