package model

type Shop struct {
	id int
	name string
}

type Order struct {
	Id     int `id`
	UserId string
	GoodId string
	ShopId string
	Number int
	Time   string
}

type Store struct {
	ShopId string
	GoodId string
	Number int
}