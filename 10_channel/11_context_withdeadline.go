package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// The variable `d` has the current time plus 11 seconds, this is the deadline
	// You can set a date as deadline
	d := time.Now().Add(11 * time.Second)

	// Here we are creting a WithDeadline context, we pass the parent context and the deadline time
	// The second argument is of type `Time`
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// It will send the end signal in case the Context doesn't reach the deadline time
	defer cancel()

	// We invoke a function with the context
	deadLine(ctx)

	// We wait 20 seconds to demostrate the deadline
	time.Sleep(20 * time.Second)
} // Here the cancel will be called, but in this example the goroutine will reach the deadline time before this point

// The function with the Context as a parameter
func deadLine(ctx context.Context) {
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
