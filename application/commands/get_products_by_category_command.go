package commands

import (
	"errors"
	"luvr.com/application/services"
	"luvr.com/domain/aggregate/product"
)

type GetProductsByCategoryCommand struct {
	productService  *services.ProductService
	categoryService *services.CategoryService
}

func (command *GetProductsByCategoryCommand) Handle(categoryId int) (*product.ProductCollection, error) {
	if command.isCategoryExist(categoryId) {
		return nil, errors.New("категории не существует")
	}

	return command.productService.GetProductsByCategoryId(categoryId), nil
}

func (command *GetProductsByCategoryCommand) isCategoryExist(categoryId int) bool {
	category := command.categoryService.GetCategoryById(categoryId)

	if category == nil {
		return false
	}

	return true
}
