package product

type UsecaseItf interface {
	GetProducts(page, pageSize int) ([]Product, error)
	GetProductByID(ID int64) (Product, error)
	CreateProduct(product Product) (Product, error)
	UpdateProduct(product Product) (Product, error)
	DeleteProduct(ID int64) error
}

type usecase struct {
	productRepo RepoItf
}

func NewProductUsecase(productRepo RepoItf) UsecaseItf {
	return &usecase{
		productRepo: productRepo,
	}
}

func (p *usecase) GetProducts(page, pageSize int) ([]Product, error) {
	return p.productRepo.GetProducts(page, pageSize)
}

func (p *usecase) GetProductByID(ID int64) (Product, error) {
	return p.productRepo.GetProductByID(ID)
}

func (p *usecase) CreateProduct(product Product) (Product, error) {
	return p.productRepo.CreateProduct(product)
}

func (p *usecase) UpdateProduct(product Product) (Product, error) {
	return p.productRepo.UpdateProduct(product)
}

func (p *usecase) DeleteProduct(ID int64) error {
	return p.productRepo.DeleteProduct(ID)
}
