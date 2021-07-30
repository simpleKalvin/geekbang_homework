package dao

type Status int32

const (
	StatusUnPay Status = iota
	StatusPay
	StatusRefund
)

type Order struct {
	id      int
	orderNo string
	money   float32
	status  Status
}

func New(goods_id int) (*Order, error) {
	// 入库操作
	return nil, nil
}
