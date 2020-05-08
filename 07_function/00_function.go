package main

import "fmt"

func main() {
	// empty function invocation with no arguments
	fmt.Println("\nCalling empty function:")
	empty()

	// hello function with one argument
	fmt.Println("\nCalling one argument function:")
	hello("one argument function")

	// register function with two arguments and two return values
	fmt.Println("\nCalling two argument and two return values:")
	message, registered := register("Karla","Martinez")
	fmt.Println(message,":",registered)

	// many arguments acting as a single variadic parameter
	fmt.Println("\nCalling many arguments into a single variadic parameter:")
	variadic(1,2,3,4,5,6,7,8,9,0)
}

// Declare function
// func (r receiver) identifier(parameters) (return(s))  { ... }
func empty()  {
	fmt.Println("Empty function")
}

// Declare function with one parameter
func hello(s string)  {
	fmt.Println("Hello,", s)
}

// Declare function with two parameters
func register(firstname string, lastname string) (string, bool) {
	message := fmt.Sprint(firstname," ",lastname, " registered")
	registered := true
	return message, registered
}

// Declare function with variadic parameter
func variadic(x ...int)  {
	fmt.Println(x) // x is a slice of int
	fmt.Printf("Type: %T\n",x)
}