package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Every 500ms"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "Every 2 seconds"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

		{
			
			var a int
			a = 4
		
		}

		/*
			Blocking on c1:
			The line msg1 := <-c1 will block until the channel c1 receives a message.
			Since you are sending to c1 every 500 milliseconds (500ms),
			it will wait for that period, then print the message.

			Blocking on c2:
			After receiving the message from c1, it will move to msg2 := <-c2.
			This will now block on c2, but since c2 only sends every 2 seconds (2s),
			it will have to wait for the full 2 seconds before printing.

			for {
				msg1 := <-c1
				msg2 := <-c2
			}

		*/

	}
}
