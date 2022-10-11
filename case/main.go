package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	data, err := getSequential(123)
	if err != nil {
		log.Printf("err %v", err)
	}
	log.Printf("data %+v", data)
	log.Printf("response time sequential %s", time.Since(start).String())

	start = time.Now()
	data, err = getConcurrent(123)
	if err != nil {
		log.Printf("err %v", err)
	}
	log.Printf("data %+v", data)
	log.Printf("response time concurrency %s", time.Since(start).String())

}

func getConcurrent(productID int64) (data ProductDisplayPage, err error) {
	var product Product
	var media Media
	var flashsale FlashSale
	var insight Insight

	var wg sync.WaitGroup
	var mtx sync.Mutex
	wg.Add(4)

	// Get Product
	go func() {
		defer wg.Done()
		var errGo error
		product, errGo = GetProduct(productID)
		if errGo != nil {
			mtx.Lock()
			err = errGo
			mtx.Unlock()
			return
		}
	}()

	// Get Media
	go func() {
		defer wg.Done()
		var errGo error
		media, errGo = GetMedia(productID)
		if err != nil {
			mtx.Lock()
			err = errGo
			mtx.Unlock()
			return
		}
	}()

	// Get Flashsale
	go func() {
		defer wg.Done()
		var errGo error
		flashsale, errGo = GetFlashsale(productID)
		if err != nil {
			mtx.Lock()
			err = errGo
			mtx.Unlock()
			return
		}
	}()

	// Get Insight
	go func() {
		defer wg.Done()
		var errGo error
		insight, errGo = GetInsight(productID)
		if err != nil {
			mtx.Lock()
			err = errGo
			mtx.Unlock()
			return
		}
	}()
	wg.Wait()

	return ProductDisplayPage{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Condition: func(conditionNum int64) string {
			if conditionNum == 1 {
				return "Baru"
			}
			return "Bekas"
		}(product.Condition),
		Price:           product.Price,
		DiscountedPrice: flashsale.DiscountedPrice,
		Rating:          insight.Rating,
		SoldStats:       insight.SoldStats,
		CountReview:     insight.CountReview,
		Discussion:      insight.Discussion,
		MainImage:       media.MainImage,
		AdditionalImage: media.AdditionalImage,
	}, nil
}

func getSequential(productID int64) (data ProductDisplayPage, err error) {
	// Get Product
	product, err := GetProduct(productID)
	if err != nil {
		return data, err
	}

	// Get Media
	media, err := GetMedia(productID)
	if err != nil {
		return data, err
	}

	// Get Flashsale
	flashsale, err := GetFlashsale(productID)
	if err != nil {
		return data, err
	}

	// Get Insight
	insight, err := GetInsight(productID)
	if err != nil {
		return data, err
	}

	return ProductDisplayPage{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Condition: func(conditionNum int64) string {
			if conditionNum == 1 {
				return "Baru"
			}
			return "Bekas"
		}(product.Condition),
		Price:           product.Price,
		DiscountedPrice: flashsale.DiscountedPrice,
		Rating:          insight.Rating,
		SoldStats:       insight.SoldStats,
		CountReview:     insight.CountReview,
		Discussion:      insight.Discussion,
		MainImage:       media.MainImage,
		AdditionalImage: media.AdditionalImage,
	}, nil
}
