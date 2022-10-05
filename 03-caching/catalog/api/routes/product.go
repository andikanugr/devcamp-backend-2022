package routes

import (
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/api/handler"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/pkg/product"
	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app fiber.Router, usecase product.UsecaseItf) {
	app.Get("/products", handler.GetProducts(usecase))
	app.Get("/products/:id", handler.GetProductByID(usecase))
	app.Post("/products", handler.CreateProduct(usecase))
	app.Put("/products", handler.UpdateProduct(usecase))
	app.Delete("/products/:id", handler.DeleteProduct(usecase))
}
