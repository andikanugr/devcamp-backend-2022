package http

import (
	"encoding/json"
	"net/http"

	entity "github.com/andikanugr/devcamp-backend-2022/clean-architecture/internal/entity/book"
)

// bookUsecase interface is defined here instead of in the usecase layer for these reasons:
// - in case of manual mock, will be easier if we only define the only methods that we need
// see https://github.com/golang/go/wiki/CodeReviewComments#interfaces
type bookUsecase interface {
	GetBooks() ([]entity.Book, error)
	Create(title string, minReaderAge int) (int64, error)
}

// Handler defines the book handler
type Handler struct {
	bookUc bookUsecase
}

// New creates book handler
func New(bookUc bookUsecase) *Handler {
	return &Handler{
		bookUc: bookUc,
	}
}

// GetBooks is handler for get books list
func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {

	// call the usecase
	book, err := h.bookUc.GetBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`error from GetBooks`))
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`failed to encode response`))
		return
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var reqBody entity.Book
	defer r.Body.Close()
	// params decode
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// call the usecase
	id, err := h.bookUc.Create(reqBody.Title, reqBody.MinReaderAge)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`error from Create`))
		return
	}
	response := map[string]int64{
		"id": id,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`failed to encode response`))
		return
	}
}

type createResponse struct {
	ID int `json:"id"`
}
