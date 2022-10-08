package product

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/andikanugr/devcamp-backend-2022/03-caching/catalog/storage"
)

type productCache struct {
	productRepo RepoItf
	redis       storage.RedisDriver
	expiration  time.Duration
}

func NewProductCache(productRepo RepoItf, redis storage.RedisDriver) RepoItf {
	return &productCache{
		productRepo: productRepo,
		redis:       redis,
		expiration:  2 * time.Minute,
	}
}

const (
	cacheKeyProduct  = "product:%d"
	cacheKeyProducts = "products:%d:%d"
)

func (p productCache) GetProducts(page, pageSize int) ([]Product, error) {
	cacheKey := fmt.Sprintf(cacheKeyProducts, page, pageSize)
	var products []Product

	val, err := p.redis.Get(context.Background(), cacheKey).Result()
	if val != "" {
		err = json.Unmarshal([]byte(val), &products)
		if err != nil {
			log.Println("func GetProducts error unmarshal products from redis, ", err.Error())
			return nil, err
		}
		return products, nil
	}

	products, err = p.productRepo.GetProducts(page, pageSize)
	if err != nil {
		log.Println("func GetProducts error when get products from db, ", err.Error())
		return nil, err
	}

	productsJSON, err := json.Marshal(products)
	if err != nil {
		log.Println("func GetProducts marshal error, ", err.Error())
		return nil, err
	}

	err = p.redis.Set(context.Background(), cacheKey, productsJSON, p.expiration).Err()
	if err != nil {
		log.Println("func GetProducts error when set redis, ", err.Error())
	}

	return products, nil
}

func (p productCache) GetProductByID(ID int64) (Product, error) {
	cacheKey := fmt.Sprintf(cacheKeyProduct, ID)
	var product Product

	val, err := p.redis.Get(context.Background(), cacheKey).Result()
	if val != "" {
		err = json.Unmarshal([]byte(val), &product)
		if err != nil {
			log.Println("func GetProductByID error unmarshal product from redis, ", err.Error())
			return product, err
		}
		return product, nil
	}

	product, err = p.productRepo.GetProductByID(ID)
	if err != nil {
		log.Println("func GetProductByID error when get product from db, ", err.Error())
		return product, err
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Println("func GetProductByID marshal error, ", err.Error())
		return product, err
	}

	err = p.redis.Set(context.Background(), cacheKey, productJSON, p.expiration).Err()
	if err != nil {
		log.Println("func GetProductByID error when set redis, ", err.Error())
	}

	return product, nil
}

func (p productCache) CreateProduct(product Product) (Product, error) {
	cacheKey := fmt.Sprintf(cacheKeyProduct, product.ID)
	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Println("func CreateProduct marshal error, ", err.Error())
		return product, err
	}

	err = p.redis.Set(context.Background(), cacheKey, productJSON, p.expiration).Err()
	if err != nil {
		log.Println("func CreateProduct error when set redis, ", err.Error())
	}

	return p.productRepo.CreateProduct(product)
}

func (p productCache) UpdateProduct(product Product) (Product, error) {
	cacheKey := fmt.Sprintf(cacheKeyProduct, product.ID)

	resp, err := p.productRepo.UpdateProduct(product)
	if err != nil {
		log.Println("func UpdateProduct error when update product to db, ", err.Error())
		return resp, err
	}

	err = p.redis.Del(context.Background(), cacheKey).Err()
	if err != nil {
		log.Println("func UpdateProduct error when delete redis, ", err.Error())
	}

	return resp, nil
}

func (p productCache) DeleteProduct(ID int64) error {
	cacheKey := fmt.Sprintf(cacheKeyProduct, ID)

	err := p.productRepo.DeleteProduct(ID)
	if err != nil {
		log.Println("func DeleteProduct error when delete product to db, ", err.Error())
		return err
	}

	err = p.redis.Del(context.Background(), cacheKey).Err()
	if err != nil {
		log.Println("func DeleteProduct error when delete redis, ", err.Error())
	}

	return nil
}
