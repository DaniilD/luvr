package commands

import (
	"errors"
	"luvr.com/application/services"
	productPkg "luvr.com/domain/aggregate/product"
)

type GetProductCommand struct {
	productService *services.ProductService
}

func (command *GetProductCommand) Handle(productId int) (*productPkg.Product, error) {
	product := command.productService.GetById(productId)

	if product == nil {
		return nil, errors.New("товара не существует")
	}

	return product, nil
}
