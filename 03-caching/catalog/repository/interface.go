package repository

import "github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/entity"

type ProductRepoItf interface {
	GetProducts() ([]entity.Product, error)
	GetProductByID(ID int64) (entity.Product, error)
	CreateProduct(product entity.Product) (entity.Product, error)
}