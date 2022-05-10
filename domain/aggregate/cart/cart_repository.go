package cart

type CartRepository interface {
	GetCartByUserId(userId int) *Cart
	AddCartItem(cartItem *CartItem)
	AddCart(cart *Cart)
	UpdateCartItem(cartItem *CartItem)
	UpdateCart(cart *Cart)
}
