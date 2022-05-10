package commands

import (
	"luvr.com/application/services"
	"luvr.com/domain/aggregate/category"
)

type GetAllCategoriesCommand struct {
	categoryService *services.CategoryService
}

func (command *GetAllCategoriesCommand) Handle() *category.CategoryCollection {
	return command.categoryService.GetAllCategories()
}
