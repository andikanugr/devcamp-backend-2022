package main

import "sync"

func main() {
	wg := &sync.WaitGroup{}

	// define # goroutine to wait
	wg.Add(2)

	go func() {
		wg.Done()
	}()

	go func() {
		wg.Done()
	}()

	// waiting for goroutine to call wg.Done()
	// x times the same as wg.Add()
	wg.Wait()
}
