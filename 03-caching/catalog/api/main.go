package main

import (
	"log"

	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/api/routes"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/pkg/product"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/storage"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	dbCfg := storage.DBConfig{
		User:     "postgres",
		Password: "postgres",
		DBName:   "devcamp_2022",
		Host:     "localhost",
		Port:     5432,
		SSLMode:  "disable",
	}
	db := storage.NewDB(dbCfg)

	redisCfg := storage.RedisConfig{
		Addr: "localhost:6379",
	}
	redis := storage.NewRedisClient(redisCfg)

	productRepo := product.NewProductRepo(db)
	productRepoWithCache := product.NewProductCache(productRepo, redis)
	productUsecase := product.NewProductUsecase(productRepoWithCache)

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Hello"))
	})
	api := app.Group("/api")
	routes.ProductRouter(api, productUsecase)

	log.Fatal(app.Listen(":3000"))
}
