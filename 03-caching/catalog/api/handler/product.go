package handler

import (
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/api/presenter"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/pkg/product"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetProducts(usecase product.UsecaseItf) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := usecase.GetProducts()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(presenter.ProductsSuccessResponse(fetched))
	}
}
