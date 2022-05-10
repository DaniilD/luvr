package product

type ProductRepository interface {
	GetById(id int) *Product
	GetByCategoryId(categoryId int) *ProductCollection
	GetByBrandId(brandId int) *ProductCollection
}
