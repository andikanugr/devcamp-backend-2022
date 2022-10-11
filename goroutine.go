package main

import (
	"fmt"
	"time"
)

func async() {
	for counter := 0; counter < 5; counter++ {
		fmt.Println(counter)
		time.Sleep(100)
	}
}

func main() {
	go async()

	for counter := 100; counter < 106; counter++ {
		fmt.Println(counter)
		time.Sleep(100)
	}
}
