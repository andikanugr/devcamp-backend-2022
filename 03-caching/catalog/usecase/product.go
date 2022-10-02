package usecase

import (
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/entity"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/repository"
)

type ProductUsecase struct {
	productRepo repository.ProductRepoItf
}

func NewProductUsecase(productRepo repository.ProductRepoItf) ProductUsecaseItf {
	return &ProductUsecase{
		productRepo: productRepo,
	}
}

func (p *ProductUsecase) GetProducts() ([]entity.Product, error) {
	return p.productRepo.GetProducts()
}

func (p *ProductUsecase) GetProductByID(ID int64) (entity.Product, error) {
	return p.productRepo.GetProductByID(ID)
}

func (p *ProductUsecase) CreateProduct(product entity.Product) (entity.Product, error) {
	return p.productRepo.CreateProduct(product)
}
