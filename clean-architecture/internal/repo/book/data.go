package book

import (
	entity "github.com/andikanugr/devcamp-backend-2022/clean-architecture/internal/entity/book"
)

var books = []entity.Book{
	{
		ID:           1,
		Title:        "The Hobbit",
		MinReaderAge: 8,
	},
}
