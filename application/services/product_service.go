package services

import (
	"luvr.com/domain/aggregate/product"
)

type ProductService struct {
	productRepository product.ProductRepository
}

func (service *ProductService) GetProductsByCategoryId(categoryId int) *product.ProductCollection {
	return service.productRepository.GetByCategoryId(categoryId)
}

func (service *ProductService) GetByBrandId(brandId int) *product.ProductCollection {
	return service.productRepository.GetByBrandId(brandId)
}

func (service *ProductService) GetById(id int) *product.Product {
	return service.productRepository.GetById(id)
}
