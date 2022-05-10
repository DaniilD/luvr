package product_item

type ProductItem struct {
	Id           int
	Name         string
	Description  string
	Composition  string
	Price        float64
	MainImageUrl string
	IsDel        bool
	BrandId      int
	CategoryId   int
}
