package main

import (
	"fmt"
)

func main() {
	// initiate channel that takes int value
	// that have capacity of 2
	var channel = make(chan int, 2)

	go func() {
		channel <- 10
		channel <- 20
	}()

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
