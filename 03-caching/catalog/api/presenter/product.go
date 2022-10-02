package presenter

import (
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/pkg/entity"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func ProductSuccessResponse(product entity.Product) *fiber.Map {
	productResp := Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
	return &fiber.Map{
		"status": true,
		"data":   productResp,
		"error":  nil,
	}
}

// ProductsSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func ProductsSuccessResponse(data []entity.Product) *fiber.Map {
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
