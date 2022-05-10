package services

import (
	"luvr.com/domain/aggregate/cart"
)

type CartService struct {
	cartRepository cart.CartRepository
	productService *ProductService
}

func (service *CartService) GetCartByUserId(userId int) *cart.Cart {
	userCart := service.cartRepository.GetCartByUserId(userId)

	if userCart == nil {
		userCart = &cart.Cart{UserId: userId}
		service.cartRepository.AddCart(userCart)
	}

	return userCart
}

func (service *CartService) AddProductInCart(userId int, productId int, count int) {
	userCart := service.GetCartByUserId(userId)
	product := service.productService.GetById(productId)

	cartItem := &cart.CartItem{
		CartId:  userCart.Id,
		Product: product,
		Count:   count,
	}

	service.cartRepository.AddCartItem(cartItem)
}

func (service *CartService) AddCart(cart *cart.Cart) {
	service.cartRepository.AddCart(cart)
}
