package location

import (
	"luvr.com/domain/entity/address"
	"luvr.com/domain/entity/city"
)

type Location struct {
	address *address.Address
	city    *city.City
}

func NewLocation(address *address.Address, city *city.City) *Location {
	return &Location{
		address: address,
		city:    city,
	}
}

func (location *Location) GetAddress() *address.Address {
	return location.address
}

func (location *Location) GetCity() *city.City {
	return location.city
}
