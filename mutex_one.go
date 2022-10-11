package main

import (
	"fmt"
	"sync"
)

func main() {
	// initiate wait group
	wg := &sync.WaitGroup{}

	// initiate empty array
	sampleArray := []int{}

	for counter := 0; counter < 20; counter++ {
		wg.Add(1)

		go func(index int) {
			sampleArray = append(sampleArray, index)
			wg.Done()
		}(counter)
	}
	wg.Wait()

	fmt.Println(sampleArray)
	fmt.Println(len(sampleArray))
}
