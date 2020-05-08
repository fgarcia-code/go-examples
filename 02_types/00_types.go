package main

import "fmt"

func main() {

	// `y` only can hold values of type `int`
	var y = 42
	// `z` only can hold values of type `string`
	var z = "This is a string"

	fmt.Printf("The variable 'y' with value '%d' has the type '%T'\n",y,y)
	fmt.Printf("The variable 'z' with value '%s' has the type '%T'\n",z,z)

}
