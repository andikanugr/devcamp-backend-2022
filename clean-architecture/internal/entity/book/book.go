package book

import (
	"fmt"
)

type Book struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	MinReaderAge int    `json:"min_reader_age"`
}

// Validate return error if book's properties are not valid
func (b *Book) Validate() error {
	switch {
	case b.Title == ``:
		return fmt.Errorf(`Title is empty`)
	case b.MinReaderAge < 0:
		return fmt.Errorf(`MinReaderAge is negative`)
	default:
		return nil
	}
}

// IsSuitableForKids returns true for any book which is appropriate for ages of 8 and less
func (b *Book) IsSuitableForKids() bool {
	return b.MinReaderAge <= 8
}
