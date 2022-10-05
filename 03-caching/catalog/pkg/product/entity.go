package product

import (
	"fmt"
)

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

func (p *Product) BuildUpdateQuery() (string, []interface{}) {
	var query string
	var args []interface{}

	var i = 1
	if p.Name != "" {
		query += fmt.Sprintf("name=$%d,", i)
		args = append(args, p.Name)
		i++
	}
	if p.Description != "" {
		query += fmt.Sprintf("description=$%d,", i)
		args = append(args, p.Description)
		i++
	}
	if p.Price != 0 {
		query += fmt.Sprintf("price=$%d,", i)
		args = append(args, p.Price)
		i++
	}
	query = fmt.Sprintf(updateProductQuery, query[:len(query)-1], p.ID)
	return query, args
}
