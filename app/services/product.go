package services

import (
	"jakmall/app/entities"
	"jakmall/app/repositories"
)

type IProductService interface {
	GetAllProduct() ([]entities.Product, error)
}

type productService struct {
	productRepository repositories.IProductRepository
}

func ProductService(productRepository repositories.IProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) GetAllProduct() ([]entities.Product, error) {
	return s.productRepository.GetAllProduct()
}
