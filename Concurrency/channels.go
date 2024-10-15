package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go count("birds", c)

	for msg := range c {
		fmt.Println(msg)
	}
	// This code is equivalent to the "range" construct when receiving from a channel.
	// Channels have a feature where we can detect whether the channel is closed.
	// For example, if we don't close the channel, the program can panic or lead to a deadlock
	// because it will enter an infinite loop. Go can detect at runtime that no further values are being sent
	// and it doesn't know when the loop will end, causing a deadlock.

	/*
	   for {
	       msg, open := <-c
	       if !open {
	           break
	       }
	       fmt.Println(msg)
	   }
	*/

	// To prevent this, the channel must be closed by the sender after all values are sent.
	// When a channel is closed, the receiver can detect this via the 'open' boolean.
	// The 'range' construct automatically handles this check, making the code cleaner.

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- fmt.Sprintf("%d %s", i, thing)
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // Close the channel when done
}
