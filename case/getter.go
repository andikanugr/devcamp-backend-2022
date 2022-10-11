package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	// client  = &http.Client{}
	baseURL = "go-routine-alpine.herokuapp.com"
)

// GetProduct GetProduct get product
func GetProduct(productID int64) (data Product, err error) {
	responseAPI, err := http.Get(fmt.Sprintf("http://%s/product/%d", baseURL, productID))
	if err != nil {
		return data, err
	}
	defer responseAPI.Body.Close()
	byteBody, err := io.ReadAll(responseAPI.Body)
	if err != nil {
		return data, err
	}
	if responseAPI.StatusCode != http.StatusOK {
		return data, fmt.Errorf("got status %d", responseAPI.StatusCode)
	}

	return data, json.Unmarshal(byteBody, &data)
}

// GetMedia GetMedia get media
func GetMedia(productID int64) (data Media, err error) {
	responseAPI, err := http.Get(fmt.Sprintf("http://%s/media/%d", baseURL, productID))
	if err != nil {
		return data, err
	}
	defer responseAPI.Body.Close()
	byteBody, err := io.ReadAll(responseAPI.Body)
	if err != nil {
		return data, err
	}
	if responseAPI.StatusCode != http.StatusOK {
		return data, fmt.Errorf("got status %d", responseAPI.StatusCode)
	}

	return data, json.Unmarshal(byteBody, &data)
}

// GetInsight GetInsight get insight
func GetInsight(productID int64) (data Insight, err error) {
	responseAPI, err := http.Get(fmt.Sprintf("http://%s/insight/%d", baseURL, productID))
	if err != nil {
		return data, err
	}
	defer responseAPI.Body.Close()
	byteBody, err := io.ReadAll(responseAPI.Body)
	if err != nil {
		return data, err
	}
	if responseAPI.StatusCode != http.StatusOK {
		return data, fmt.Errorf("got status %d", responseAPI.StatusCode)
	}

	return data, json.Unmarshal(byteBody, &data)
}

// GetFlashsale GetFlashsale get flashsale
func GetFlashsale(productID int64) (data FlashSale, err error) {
	responseAPI, err := http.Get(fmt.Sprintf("http://%s/flashsale/%d", baseURL, productID))
	if err != nil {
		return data, err
	}
	defer responseAPI.Body.Close()
	byteBody, err := io.ReadAll(responseAPI.Body)
	if err != nil {
		return data, err
	}
	if responseAPI.StatusCode != http.StatusOK {
		return data, fmt.Errorf("got status %d", responseAPI.StatusCode)
	}

	return data, json.Unmarshal(byteBody, &data)
}
