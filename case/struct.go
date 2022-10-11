package main

type (
	ProductDisplayPage struct {
		ID              int64   `json:"id"`
		Name            string  `json:"name"`
		Description     string  `json:"description"`
		Condition       string  `json:"condition"`
		Price           int64   `json:"price"`
		DiscountedPrice int64   `json:"discounted_price"`
		Rating          float64 `json:"rating"`
		SoldStats       int64   `json:"sold_stats"`
		CountReview     int64   `json:"count_review"`
		Discussion      int64   `json:"discussion"`
		ResponseTime    string  `json:"response_time"`
		MainImage       Image   `json:"main_image"`
		AdditionalImage []Image `json:"additional_image"`
	}

	Product struct {
		ID           int64  `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Condition    int64  `json:"condition"`
		Price        int64  `json:"price"`
		ResponseTime string `json:"response_time"`
	}

	FlashSale struct {
		ID              int64  `json:"id"`
		ProductID       int64  `json:"product_id"`
		DiscountedPrice int64  `json:"discounted_price"`
		IsActive        bool   `json:"is_active"`
		ResponseTime    string `json:"response_time"`
	}

	Media struct {
		ProductID       int64   `json:"product_id"`
		MainImage       Image   `json:"main_image"`
		AdditionalImage []Image `json:"additional_image"`
		ResponseTime    string  `json:"response_time"`
	}

	Image struct {
		ID           int64  `json:"id"`
		FilePath     string `json:"file_path"`
		FileName     string `json:"file_name"`
		BaseURL      string `json:"base_url"`
		ResponseTime string `json:"response_time"`
	}

	Insight struct {
		ProductID    int64   `json:"product_id"`
		Rating       float64 `json:"rating"`
		SoldStats    int64   `json:"sold_stats"`
		CountReview  int64   `json:"count_review"`
		Discussion   int64   `json:"discussion"`
		ResponseTime string  `json:"response_time"`
	}
)
