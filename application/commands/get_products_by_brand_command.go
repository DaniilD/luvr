package commands

import (
	"errors"
	"luvr.com/application/services"
	"luvr.com/domain/aggregate/product"
)

type GetProductsByBrandCommand struct {
	productService *services.ProductService
	brandService   *services.BrandService
}

func (command *GetProductsByBrandCommand) Handle(brandId int) (*product.ProductCollection, error) {
	if !command.isBrandExist(brandId) {
		return nil, errors.New("бренда не существует")
	}

	return command.productService.GetByBrandId(brandId), nil
}

func (command *GetProductsByBrandCommand) isBrandExist(brandId int) bool {
	brand := command.brandService.GetBrandById(brandId)

	if brand == nil {
		return false
	}

	return true
}
