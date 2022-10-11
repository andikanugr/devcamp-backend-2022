package main

import (
	"math/rand"
	"time"

	bookhandler "github.com/andikanugr/devcamp-backend-2022/clean-architecture/internal/handler/http/book"
	bookrepo "github.com/andikanugr/devcamp-backend-2022/clean-architecture/internal/repo/book"
	bookuc "github.com/andikanugr/devcamp-backend-2022/clean-architecture/internal/usecase/book"
)

func startApp() error {

	// connect the db
	// connect redis
	// tracer with jaeger
	// init dependancies
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// DI Repo
	bookRepo, err := bookrepo.New(random)
	if err != nil {
		return err
	}
	// DI Usecase
	bookUc := bookuc.New(bookRepo)

	// DI Handler
	bookHandler := bookhandler.New(bookUc)

	router := newRoutes(bookHandler)
	return startServer(router)
}
