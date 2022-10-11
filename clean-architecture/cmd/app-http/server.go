package main

import (
	"fmt"
	"net/http"
	"time"
)

func startServer(handler http.Handler) error {
	srv := http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      handler,
		Addr:         ":8080",
	}

	fmt.Println("server is running on port 8080")
	return srv.ListenAndServe()
}
