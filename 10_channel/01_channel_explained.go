package main

import (
	"fmt"
	"time"
)

func main() {
	// This creates a channel
	var c = make(chan int)

	// It invokes a new goroutine
	go func() {
		// This is the second print, it will print because nothing block it, it is launched concurrently
		fmt.Println("I'm in the channelFunction that writes into the channel\n")

		// This will block the main goroutine at line 31, because this line will wait for the send operation at line 21
		time.Sleep(10 * time.Second)

		// It will send the value 10 through the channel releasing the line 31 in the main goroutine
		c <- 10

		// The goroutine will continue with its normal flow
		fmt.Println("I'm in the channelFunction, I've written in the channel, also the channel have been read\n")
	}()

	// This is the first print, it will print because nothing block it
	fmt.Println("I'm in main goroutine\n")

	// The expression will wait until another goroutine send data through the channel, then the value will be printed
	fmt.Println("The value in the channel is:", <-c, "\n") // The channel blocks

	// This print will be executed after the channel has received the data
	fmt.Println("I'm in a main goroutine and I've read the channel\n")
}
