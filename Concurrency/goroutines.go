package main

import (
	"fmt"
	"time"
)

func main() {
	go count("birds")
	count("cars")
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
