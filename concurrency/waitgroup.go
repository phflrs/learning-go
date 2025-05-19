/*

The sync.WaitGroup in Go is a struct used to wait for a collection of goroutines to finish. It operates as a counter that is incremented by the number of goroutines to wait for, and decremented each time a goroutine completes its execution. The Wait method blocks until the counter becomes zero, indicating that all goroutines have finished.
The sync.WaitGroup has three primary methods:
- Add(delta int):
		Adds delta, which can be positive or negative, to the WaitGroup counter. It is typically called before starting a goroutine to indicate that there's one more task to wait for.
- Done():
		Decrements the WaitGroup counter by one. It is usually called at the end of a goroutine's execution to signal that it has completed. It's equivalent to Add(-1).
- Wait():
		Blocks the execution of the current goroutine until the WaitGroup counter becomes zero.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter before starting the goroutine
		go worker(i, &wg)
	}

	wg.Wait() // Block until the counter is zero
	fmt.Println("All workers done")
}