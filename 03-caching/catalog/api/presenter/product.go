package presenter

import (
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/pkg/product"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

func ProductSuccessResponse(product product.Product) *fiber.Map {
	productResp := Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
	return &fiber.Map{
		"status": true,
		"data":   productResp,
		"error":  nil,
	}
}

// ProductsSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func ProductsSuccessResponse(data []product.Product) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// ProductErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ProductErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
