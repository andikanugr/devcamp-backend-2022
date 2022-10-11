package book

import (
	entity "github.com/andikanugr/devcamp-backend-2022/clean-architecture/internal/entity/book"
)

type bookRepo interface {
	GetBooks() ([]entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
}

type Usecase struct {
	bookRepo bookRepo
}

func New(books bookRepo) *Usecase {
	return &Usecase{
		bookRepo: books,
	}
}

func (uc *Usecase) GetBooks() ([]entity.Book, error) {
	return uc.bookRepo.GetBooks()
}

func (uc *Usecase) Create(title string, minReaderAge int) (int64, error) {
	book := entity.Book{
		Title:        title,
		MinReaderAge: minReaderAge,
	}
	if err := book.Validate(); err != nil {
		return 0, err
	}

	savedBook, err := uc.bookRepo.Create(book)
	if err != nil {
		return 0, err
	}

	return savedBook.ID, nil
}
