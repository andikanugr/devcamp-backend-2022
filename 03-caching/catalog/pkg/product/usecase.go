package product

import (
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/pkg/entity"
)

type UsecaseItf interface {
	GetProducts() ([]entity.Product, error)
	GetProductByID(ID int64) (entity.Product, error)
	CreateProduct(product entity.Product) (entity.Product, error)
}

type usecase struct {
	productRepo RepoItf
}

func NewProductUsecase(productRepo RepoItf) UsecaseItf {
	return &usecase{
		productRepo: productRepo,
	}
}

func (p *usecase) GetProducts() ([]entity.Product, error) {
	return p.productRepo.GetProducts()
}

func (p *usecase) GetProductByID(ID int64) (entity.Product, error) {
	return p.productRepo.GetProductByID(ID)
}

func (p *usecase) CreateProduct(product entity.Product) (entity.Product, error) {
	return p.productRepo.CreateProduct(product)
}
