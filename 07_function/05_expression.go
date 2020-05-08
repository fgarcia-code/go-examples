package main

import "fmt"

func main() {
	f := func(message string){ // function as a type
		fmt.Println(message,"function expression")
	}

	f("Hello") // Calling the function through a variable
}
