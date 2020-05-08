package main

import "fmt"

func main() {
	sum := operation(sum,1,2,3,4) // Invocation of callback function with Function 1
	product := operation(product,1,2,3,4) // Invocation of callback function with Function 2
	fmt.Println("The result of the sum is:",sum)
	fmt.Println("The result of the product is:",product)
}

func operation(o func(p ...int)int, values ...int) int  { // Callback function func(...int) as parameter
	return o(values...)
}

func sum(values ...int) int  { // Function 1 of type func(...int) int
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func product(values ...int) int  { // Function 2 of type func(...int) int
	product := 1
	for _, value := range values {
		product *= value
	}
	return product
}