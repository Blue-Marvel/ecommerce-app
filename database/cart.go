package database

import "errors"

var (
	ErrCantFindProduct        = errors.New("can't find product")
	ErrCantDecodeProduct      = errors.New("can't decode product")
	ErrUserIdIsNotValid       = errors.New("user id is not valid")
	ErrCantUpdateUser         = errors.New("can't add this product to cart")
	ErrCantRemoveItemFromCart = errors.New("can't remove item from cart")
	ErrCantGetItemFromCart    = errors.New("can't get item from cart")
	ErrCantBuyCartItem        = errors.New("can't buy cart item")
)

func AddProductToCart() {}

func RemoveProductFromCart() {}

func BuyFromCart() {}

func InstantBuy() {}
