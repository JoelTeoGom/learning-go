package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // Add one task to wait for

	go func() {
		count("birds")
		wg.Done() // Indicate the goroutine is finished
	}()

	wg.Wait() // Wait for all goroutines to complete
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
