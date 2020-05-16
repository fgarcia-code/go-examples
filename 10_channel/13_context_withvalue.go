package main

import (
	"context"
	"fmt"
	"time"
)

// You must define your own key type to avoid collisions whit other packages
type favContextKey string

func main() {
	// This is the key of the context
	key := favContextKey("language")

	// We create a new WithValue Context, with a key and a value
	ctx := context.WithValue(context.Background(), key, "Golang")

	// We pass the context and the key to two functions as argument
	first(ctx,key)
	second(ctx,key)

	// We wait untill all finish
	time.Sleep(10 * time.Second)
}

// We define a function with a context and a key as parameter
func first(ctx context.Context, key favContextKey) {
	go func() {
		fmt.Println("First function")
		// We print the value of the context
		fmt.Println(ctx.Value(key))
	}()
}

// We define a function with a context and a key as parameter
func second(ctx context.Context, key favContextKey) {
	go func() {
		fmt.Println("Second function")
		// We print the value of the context
		fmt.Println(ctx.Value(key))
	}()
}
