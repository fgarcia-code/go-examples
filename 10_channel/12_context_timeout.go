package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// This is similar to WithDeadline, the difference is the second argument it is of type `Duration`
	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Second)

	// It will send the end signal in case the Context doesn't reach the timeout time
	defer cancel()

	// We invoke a function with the context
	timeout(ctx)

	// We wait 20 seconds to demostrate the timeout
	time.Sleep(20 * time.Second)
} // Here the cancel will be called, but in this example the goroutine will reach the timeout time before this point

// The function with the Context as a parameter
func timeout(ctx context.Context) {
	// The goroutine
	go func() {
		// This for will never end until a condition inside it
		for {
			select {
			case t := <-time.After(1 * time.Second):
				// It will be executed every second
				fmt.Println("The time is:", t)
			case <-ctx.Done():
				// It will be called when we reach the deadline time, then it will finish the goroutine
				fmt.Println(ctx.Err())
				return
			}
		}
	}()
}
