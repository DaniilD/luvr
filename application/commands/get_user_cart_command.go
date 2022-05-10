package commands

import (
	"errors"
	"luvr.com/application/services"
	"luvr.com/domain/aggregate/cart"
)

type GetUserCartCommand struct {
	cartService *services.CartService
}

func (command *GetUserCartCommand) Handle(userId int) (*cart.Cart, error) {

	if !command.isValidUser(userId) {
		return nil, errors.New("пользователя не существует")
	}

	return command.cartService.GetCartByUserId(userId), nil
}

func (command *GetUserCartCommand) isValidUser(userId int) bool {
	return true
}