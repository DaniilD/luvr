package cart

import "luvr.com/domain/aggregate/product"

type CartItem struct {
	Id      int
	CartId  int
	Product *product.Product
	Count   int
}
