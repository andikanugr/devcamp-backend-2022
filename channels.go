package main

import (
	"fmt"
)

func main() {
	// initiate channel that takes int value
	var channel = make(chan int)

	go func() {
		channel <- 10
	}()

	fmt.Println(<-channel)
}
