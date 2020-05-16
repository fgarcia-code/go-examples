package main

import (
	"context"
	"fmt"
)

func main() {
	// It will create an emptyCtx parent context, the emptyCtx implements the Context methods
	// The methods in the emptyCtx will always return `nil`
	bgctx := context.Background()

	// It will create a Context of type `cancelCtx`, one of its fields is a Context (the parent Context, background or TODO)
	// This new Context has new fields including the parent Context, and it implements the Context methods
	// One of the methods it implements, is the `Done()` method, this method will return a channel
	// You can use the channel returned by `Done()` method to cancel a goroutine in conjunction with the `cancel` function
	ctx, cancel := context.WithCancel(bgctx)

	// It will invoke the function `cancel` at the end of the main goroutine
	// It will send the end signal to the context, terminating all the goroutines that share the context `ctx`
	defer cancel()

	// The function `generatorInts()` will return a channel and the for will iterate over this channel
	// We pass the context as an argument, only the goroutine inside `generatorInts()` function will be in the `ctx` context
	for n := range generatorInts(ctx) {
		fmt.Println(n)
		// It will only retrieve 10 elements from the channel, then it will exit
		if n == 10 {
			break
		}
	}
} // Here the cancel will be called and the goroutine inside the `generatorInts()` will be finished

// The `generatorInts()` function with a channel as a return, and the context as a parameter
func generatorInts(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	// This goroutine will be in the scope of the context `ctx`
	go func() {
		// This for will never end until a condition inside it
		for {
			// This select has two cases
			select {
			case <-ctx.Done():
				// In this case the `Done()` method will be called when the main goroutine calls the `cancel()` function
				// This case only will be executed when it receives a value through the channel returned by `Done()`
				// This return will finish the current goroutine
				return
			case dst <- n:
				// This case will be executed every time the channel is available to send data through the `dst` channel
				// Then it will be increment by 1 the variable n
				n++
			}
		}
	}()

	// This returns the `dst` channel and it will be consumed by the for at line 25
	return dst
}
