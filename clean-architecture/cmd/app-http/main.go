package main

import (
	"log"
)

func main() {
	// init config
	// init logger
	err := startApp()
	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
