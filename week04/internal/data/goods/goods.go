package goods

import "fmt"

type GoodsId int64

type Goods struct {
	Id GoodsId
	Name string
	Price float64
}



func NewGoods(goodsId GoodsId) Goods  {
	return Goods{
		Id:    goodsId,
		Name:  fmt.Sprintf("%s_%v", "goods", goodsId),
		Price: 12.1,
	}
}