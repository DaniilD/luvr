package brand

type BrandCollection struct {
	brands []*Brand
}

func (collection *BrandCollection) Add(brand *Brand) {
	collection.brands = append(collection.brands, brand)
}
