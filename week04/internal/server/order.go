package server

import (
	"context"
	v1 "github.com/simpleKalvin/geekbang_homework/week04/api/order/v1"
	"github.com/simpleKalvin/geekbang_homework/week04/internal/biz"
)

type OrderServer struct {
	v1.UnimplementedOrderServer
	order *biz.Order
}

func NewOrderServer(user_id, goods_id int64,) *OrderServer {
	return &OrderServer{}
}

func (this OrderServer) Order(ctx context.Context, in *v1.OrderRequest) (*v1.OrderReply, error) {
	return &v1.OrderReply{}, nil
}