package category

type CategoryRepository interface {
	GetAllCategories() *CategoryCollection
	GetById() *Category
}
