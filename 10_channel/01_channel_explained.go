package main

import (
	"fmt"
	"sync"
	"time"
)

// This creates a `Channel`
var c = make(chan int)

// We create a `WaitGroup` only to show the program's flow through prints
var wg sync.WaitGroup

func main() {
	wg.Add(1) // Ignore it

	// It invokes a new goroutine
	go channelFunction()

	// This is the first `Print`, it will print because nothing lock it
	fmt.Println("I'm in main goroutine\n")

	// The expression will wait until another goroutine send data through the `Channel`, then the value will be printed
	fmt.Println("The value in the channel is:", <-c,"\n")

	// This `Print` will be executed after the `Channel` have been received the data
	fmt.Println("I'm in a main goroutine and I've read the channel\n")

	wg.Wait() // Ignore it
}

func channelFunction() {
	defer wg.Done() // Ignore it

	// This is the second `Print`, it will print because nothing block it
	fmt.Println("I'm in the channelFunction that writes into the channel\n")

	// This will lock the `main` goroutine at line 25, because it will wait for the send operation at line 44
	time.Sleep(10 * time.Second)

	// It will send the value 10 through the `Channel` unlocking the line 25 in the `main` goroutine
	// This goroutine will be locked at this line, until someone else receive the value from the `Channel`
	c <- 10

	// The goroutine will continue with its normal flow
	fmt.Println("I'm in the channelFunction, I've written in the channel, also the channel have been read\n")
}