package handler

import (
	"net/http"

	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/api/presenter"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/pkg/product"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/util/converter"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(usecase product.UsecaseItf) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page := c.Query("page")
		pageSize := c.Query("page_size")

		fetched, err := usecase.GetProducts(converter.ToInt(page), converter.ToInt(pageSize))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(presenter.ProductsSuccessResponse(fetched))
	}
}

func GetProductByID(usecase product.UsecaseItf) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID := c.Params("id")
		fetched, err := usecase.GetProductByID(converter.ToInt64(ID))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(presenter.ProductSuccessResponse(fetched))
	}
}

func CreateProduct(usecase product.UsecaseItf) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var product product.Product
		if err := c.BodyParser(&product); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		fetched, err := usecase.CreateProduct(product)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(presenter.ProductSuccessResponse(fetched))
	}
}

func UpdateProduct(usecase product.UsecaseItf) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var product product.Product
		if err := c.BodyParser(&product); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		fetched, err := usecase.UpdateProduct(product)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(presenter.ProductSuccessResponse(fetched))
	}
}

func DeleteProduct(usecase product.UsecaseItf) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID := c.Params("id")
		err := usecase.DeleteProduct(converter.ToInt64(ID))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(presenter.ProductSuccessResponse(product.Product{}))
	}
}
