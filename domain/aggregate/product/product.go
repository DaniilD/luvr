package product

import (
	"luvr.com/domain/aggregate/category"
	"luvr.com/domain/entity/brand"
	"luvr.com/domain/entity/gallery_item"
	"luvr.com/domain/entity/product_item"
)

type Product struct {
	product      *product_item.ProductItem
	brand        *brand.Brand
	galleryItems []*gallery_item.GalleryItem
	category     *category.Category
}

func NewProduct(product *product_item.ProductItem, brand *brand.Brand, galleryItem []*gallery_item.GalleryItem) *Product {
	return &Product{
		product:      product,
		brand:        brand,
		galleryItems: galleryItem,
	}
}

func (product *Product) GetProduct() *product_item.ProductItem {
	return product.product
}

func (product *Product) GetBrand() *brand.Brand {
	return product.brand
}

func (product *Product) GetGalleryItems() []*gallery_item.GalleryItem {
	return product.galleryItems
}
