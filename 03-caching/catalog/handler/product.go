package handler

import (
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/usecase"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ProductUC usecase.ProductUsecaseItf
}

func NewProduct(productUC usecase.ProductUsecaseItf) *Product {
	return &Product{
		ProductUC: productUC,
	}
}

func (p *Product) GetProducts(c *fiber.Ctx) error {
	products, err := p.ProductUC.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(products)
}
