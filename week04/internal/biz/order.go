package biz

import (
	"fmt"
	pb "github.com/simpleKalvin/geekbang_homework/week04/api/order/v1"
	"github.com/simpleKalvin/geekbang_homework/week04/internal/data/goods"
	"github.com/simpleKalvin/geekbang_homework/week04/internal/data/order"
	"github.com/simpleKalvin/geekbang_homework/week04/internal/data/user"
	"math/rand"
	"time"
)

type Order struct {
	order *pb.OrderReply
	pb.UnimplementedOrderServer
}

func NewOrder(user user.User, goods goods.Goods) Order {
	orderId := fmt.Sprintf("SDTD%v%v", time.Now(), rand.Int())
	orders := order.Order{
		Id:         orderId,
		UserId:     int(user.Id),
		GoodsId:    int(goods.Id),
		Price:      goods.Price,
		CreateTime: time.Now().Unix(),
	}
	fmt.Println(orders)
	orderResponse := &pb.OrderReply{OrderId: orderId}
	return Order{
		order: orderResponse,
	}
}
