package services

import "luvr.com/domain/entity/brand"

type BrandService struct {
	brandRepository brand.BrandRepository
}

func (service *BrandService) GetAllBrands() *brand.BrandCollection {
	return service.brandRepository.GetAllBrands()
}

func (service *BrandService) GetBrandById(brandId int) *brand.Brand {
	return service.brandRepository.GetById(brandId)
}
