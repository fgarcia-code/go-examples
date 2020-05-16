// This code will not run
package main

import "fmt"

func main() {
	// This creates a `Channel`
	c := make(chan int)

	// We pass the value 10 through the `Channel`
	// The `Channel` locks the goroutine at this line until another goroutine send some data through the `Channel`
	c <- 10

	// This will not run because the `Channel` at line 12, locks the main goroutine,
	// This causes a fatal error (deadlock)
	fmt.Println(<-c)
}