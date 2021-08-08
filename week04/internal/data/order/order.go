package order

type Order struct {
	Id string
	UserId int
	GoodsId int
	Price float64
	CreateTime int64
}