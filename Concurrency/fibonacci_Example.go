package main
package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start 4 worker goroutines
	for w := 1; w <= 4; w++ {
		go worker(w, jobs, results)
	}

	// Send 10 jobs to the jobs channel
	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs) // Close the jobs channel

	// Collect the results
	for r := 0; r < 10; r++ {
		fmt.Println(<-results)
	}
}

// Worker function that calculates Fibonacci for a given job
func worker(id int, jobs <-chan int, results chan<- int) {
	for n := range jobs {
		fmt.Printf("Worker %d calculating fib(%d)\n", id, n)
		results <- fib(n)
	}
}

// Recursive Fibonacci calculation
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
