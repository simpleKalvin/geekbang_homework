//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/simpleKalvin/geekbang_homework/week04/internal/biz"
	"github.com/simpleKalvin/geekbang_homework/week04/internal/data/goods"
	"github.com/simpleKalvin/geekbang_homework/week04/internal/data/user"
)


func InitOrder(goods_id goods.GoodsId, user_id user.UserId) (biz.Order, error) {
	panic(wire.Build(user.NewUser, goods.NewGoods, biz.NewOrder))
	return biz.Order{}, nil
}