package main

import (
	"fmt"
	"os"
)

func main() {
	env := "LOCAL"
	if value, ok := os.LookupEnv("ENV"); ok {
		env = value
	}
	fmt.Println("Hello Tokopedia Devcamp 2022 participants!")
	fmt.Printf("This message is run on %s environment\n", env)
}
