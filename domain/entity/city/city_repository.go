package city

type CityRepository interface {
	GetByIds(ids []int) []*City
	AddCity(city *City) int
	UpdateCity(city *City)
}
