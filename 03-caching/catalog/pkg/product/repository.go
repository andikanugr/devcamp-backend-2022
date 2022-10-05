package product

import (
	"database/sql"
	"log"
)

type RepoItf interface {
	GetProducts(page, pageSize int) ([]Product, error)
	GetProductByID(ID int64) (Product, error)
	CreateProduct(product Product) (Product, error)
	UpdateProduct(product Product) (Product, error)
	DeleteProduct(ID int64) error
}

type repo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) RepoItf {
	return &repo{
		db: db,
	}
}

func (p *repo) GetProducts(page, pageSize int) ([]Product, error) {
	limit, offset := getLimitOffset(page, pageSize)

	rows, err := p.db.Query(selectProductQuery, limit, offset)
	if err != nil {
		log.Println("func GetProducts error, ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			log.Println("func GetProducts error scanning, ", err.Error())
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *repo) GetProductByID(ID int64) (Product, error) {
	var product Product
	err := p.db.QueryRow(selectProductByIDQuery, ID).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
	)
	if err != nil {
		log.Println("func GetProductByID error, ", err.Error())
		return product, err
	}
	return product, nil
}

func (p *repo) CreateProduct(product Product) (Product, error) {
	var id int64
	err := p.db.QueryRow(createProductByIDQuery, product.Name, product.Description, product.Price).Scan(&id)
	if err != nil {
		log.Println("func CreateProduct error exec, ", err.Error())
		return product, err
	}

	product.ID = id
	return product, nil
}

func (p *repo) UpdateProduct(product Product) (Product, error) {
	updateQuery, fields := product.BuildUpdateQuery()
	resp, err := p.db.Exec(updateQuery, fields...)
	if err != nil {
		log.Println("func updateProductQuery error exec, ", err.Error())
		return product, err
	}
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		log.Println("func updateProductQuery error getting rows affected, ", err.Error())
		return product, err
	}
	if rowsAffected == 0 {
		return product, sql.ErrNoRows
	}

	return product, nil
}

func (p *repo) DeleteProduct(ID int64) error {
	_, err := p.db.Exec(deleteProductByIDQuery, ID)
	if err != nil {
		log.Println("func DeleteProduct error exec, ", err.Error())
		return err
	}
	return nil
}

func getLimitOffset(page, pageSize int) (int, int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 5
	}
	limit := pageSize
	offset := (page - 1) * pageSize
	return limit, offset
}
