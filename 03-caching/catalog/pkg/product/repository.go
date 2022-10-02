package product

import (
	"database/sql"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/pkg/entity"
	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/storage"
	"log"
)

type RepoItf interface {
	GetProducts() ([]entity.Product, error)
	GetProductByID(ID int64) (entity.Product, error)
	CreateProduct(product entity.Product) (entity.Product, error)
}

type repo struct {
	db    *sql.DB
	redis storage.RedisDriver
}

func NewProductRepo(db *sql.DB, redis storage.RedisDriver) RepoItf {
	return &repo{
		db:    db,
		redis: redis,
	}
}

func (p *repo) GetProducts() ([]entity.Product, error) {
	sql := "SELECT * FROM product"

	rows, err := p.db.Query(sql)
	if err != nil {
		log.Println("[repository][GetProducts] error, ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			log.Println("[repository][GetProducts] error scanning, ", err.Error())
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *repo) GetProductByID(ID int64) (entity.Product, error) {
	query := "SELECT * FROM product WHERE id = ?"

	var product entity.Product
	err := p.db.QueryRow(query, ID).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		log.Println("[repository][GetProductByID] error, ", err.Error())
		return product, err
	}
	return product, nil
}

func (p *repo) CreateProduct(product entity.Product) (entity.Product, error) {
	query := "INSERT INTO product (name, price) VALUES (?, ?)"

	result, err := p.db.Exec(query, product.Name, product.Price)
	if err != nil {
		log.Println("[repository][CreateProduct] error, ", err.Error())
		return product, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Println("[repository][CreateProduct] error, ", err.Error())
		return product, err
	}

	product.ID = lastID
	return product, nil
}
