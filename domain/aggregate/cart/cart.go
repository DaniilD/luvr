package cart

type Cart struct {
	Id       int
	UserId   int
	Products []*CartItem
}
