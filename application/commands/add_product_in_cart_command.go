package commands

import (
	"errors"
	"luvr.com/application/services"
)

type AddProductInCartCommand struct {
	productService *services.ProductService
	cartService    *services.CartService
}

func (command *AddProductInCartCommand) Handle(userId int, productId int, count int) error {

	if !command.isValidUser(userId) {
		return errors.New("пользователя не существует")
	}

	if !command.isValidProduct(productId) {
		return errors.New("товар не найден")
	}

	command.cartService.AddProductInCart(userId, productId, count)

	return nil
}

func (command *AddProductInCartCommand) isValidUser(userId int) bool {
	return true
}

func (command *AddProductInCartCommand) isValidProduct(productId int) bool {
	product := command.productService.GetById(productId)

	if product == nil {
		return false
	}

	return true
}
