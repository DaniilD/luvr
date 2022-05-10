package address

type AddressRepository interface {
	GetByIds(ids []int) []*Address
	AddAddress(address *Address) int
	UpdateAddress(address *Address)
}
