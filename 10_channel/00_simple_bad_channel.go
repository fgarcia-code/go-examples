// This code will not run
package main

import "fmt"

func main() {
	// This creates a channel
	c := make(chan int)

	// We pass the value 10 through the channel
	// The channel blocks the goroutine at this line until another goroutine send some data through the channel
	c <- 10

	// This will not run because the channel at line 12, blocks the main goroutine,
	// This causes a fatal error (deadlock)
	fmt.Println(<-c)
}
