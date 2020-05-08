package main

import "fmt"

func main() {
	var i interface{}
	x := 10.15
	i = x
	// Type switch
	fmt.Println("\nType switch:")
	switch  i.(type) {
	case nil:
		fmt.Println("x is nil\n")
	case int:
		fmt.Println("x is int\n")
	case float64:
		fmt.Println("x is float\n")
	case bool:
		fmt.Println("type is bool\n")
	case string:
		fmt.Println("type is string\n")
	default:
		fmt.Println("don't know the type\n")
	}
}
