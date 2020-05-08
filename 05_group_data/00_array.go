package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println("'x' is the array: ",x)
	x[3] = 15
	fmt.Println("'x[3]' is ",x)
	fmt.Println("The lenght is:",len(x))

}