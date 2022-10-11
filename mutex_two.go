package main

import (
	"fmt"
	"sync"
)

func main() {
	// initiate wait group
	wg := &sync.WaitGroup{}

	// initiate mutex
	mtx := sync.Mutex{}

	// initiate empty array
	sampleArray := []int{}

	for counter := 0; counter < 20; counter++ {
		wg.Add(1)

		go func(index int) {
			mtx.Lock()
			sampleArray = append(sampleArray, index)
			mtx.Unlock()
			wg.Done()
		}(counter)
	}
	wg.Wait()

	fmt.Println(sampleArray)
	fmt.Println(len(sampleArray))
}
