package main

import (
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/handler"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/repository"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/storage"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/usecase"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	dbCfg := storage.DBConfig{
		User:     "postgres",
		Password: "postgres",
		DBName:   "devcamp_2022",
		Host:     "localhost",
		Port:     5432,
		SSLMode:  "disable",
	}
	db := storage.NewDB(dbCfg)

	redisCfg := storage.RedisConfig{Addr: "localhost:6379"}
	redis := storage.NewRedisClient(redisCfg)

	productRepo := repository.NewProductRepo(db, redis)
	productUC := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProduct(productUC)

	app.Get("/products", productHandler.GetProducts)

	app.Listen(":3000")
}
