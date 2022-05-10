package services

import "luvr.com/domain/aggregate/category"

type CategoryService struct {
	categoryRepository category.CategoryRepository
}

func (service *CategoryService) GetAllCategories() *category.CategoryCollection {
	return service.categoryRepository.GetAllCategories()
}

func (service *CategoryService) GetCategoryById(categoryId int) *category.Category {
	return service.categoryRepository.GetById()
}
