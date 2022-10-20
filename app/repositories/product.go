package repositories

import (
	"encoding/json"
	"io/ioutil"
	"jakmall/app/entities"
	"os"
)

type IProductRepository interface {
	GetAllProduct() ([]entities.Product, error)
}

type productRepository struct {
}

func ProductRepository() *productRepository {
	return &productRepository{}
}

func (r *productRepository) GetAllProduct() ([]entities.Product, error) {
	jsonFile, err := os.Open("./resource/products.json")
	var products []entities.Product

	if err != nil {
		return products, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &products)

	return products, nil
}
