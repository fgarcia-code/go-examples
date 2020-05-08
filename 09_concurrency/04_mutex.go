// Run using go run -race 04_mutex.go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Prints the machine information
	fmt.Println("#CPU:", runtime.NumCPU())
	fmt.Println("#Goroutines:", runtime.NumGoroutine())

	// It will be a global variable for all the goroutines (Communicate by sharing)
	global_count := 0

	// The number of goroutines that will be executed
	const subroutines = 100

	// The `WaitGroup` that will take control over the goroutines
	var wg sync.WaitGroup

	// We add the number of goroutines we want the `WaitGroup` will wait for
	wg.Add(subroutines)

	// Creating a Mutex
	var mutex sync.Mutex

	// This loop will create all the goroutines
	for i := 0; i < subroutines; i++ {
		// The goroutine execution as anonymous function inside the loop
		go func() {
			// The scope of `mutex` is the main function, it will lock this block of code
			// This block of code takes precedence over the other goroutines because of `mutex` lock
			mutex.Lock()

			// The `local_count` scope only exists in the anonymous function
			// `global_count` is shared across all goroutines
			local_count := global_count // "Communicating by sharing"

			runtime.Gosched() // It doesn't works because of Mutex lock

			// The goroutine is resumed automatically here
			local_count++

			// "Communicating by sharing"
			global_count = local_count // Global count is modified, just by one goroutine at a time

			// It releases the lock allowing another goroutine to take the `mutex` lock
			mutex.Unlock()

			// The method `Done` decrements the `WaitGroup` counter, then the anonymous function exits
			wg.Done()
		}()
	}

	// Blocks the main goroutine until the `WaitGroup` counter is 0
	wg.Wait()

	// It prints only 1 goroutine (main) because all the anonymous goroutines exit before the call to `Wait()`
	fmt.Println("#Goroutines:", runtime.NumGoroutine())

	// There's no data race because the Mutex avoid it
	fmt.Println("No data race:", global_count)
}
