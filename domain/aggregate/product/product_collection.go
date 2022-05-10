package product

type ProductCollection struct {
	products []*Product
}

func (collection *ProductCollection) Add(product *Product) {
	collection.products = append(collection.products, product)
}
