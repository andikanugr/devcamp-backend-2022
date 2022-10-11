package book

import (
	entity "github.com/andikanugr/devcamp-backend-2022/clean-architecture/internal/entity/book"
)

type rng interface {
	Int63() int64
}

type Repo struct {
	rng rng
}

func New(rng rng) (*Repo, error) {
	return &Repo{
		rng: rng,
	}, nil
}

func (r *Repo) GetBooks() ([]entity.Book, error) {
	return books, nil
}

func (r *Repo) Create(book entity.Book) (entity.Book, error) {
	book.ID = r.rng.Int63()
	books = append(books, book)

	return book, nil
}
