package brand

type BrandRepository interface {
	GetAllBrands() *BrandCollection
	GetById(brandId int) *Brand
}
