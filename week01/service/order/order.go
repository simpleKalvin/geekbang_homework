package order

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/simpleKalvin/geekbang_homework/week01/dao"
)

func New(id int) (*dao.Order, error) {
	/**
	 * @Description:
	1. 根据id查询产品
	2. 产品存在则创建订单
	3. 产品不存在则不创建订单并返回异常
	 */
	goodsData, err := dao.GetOne(id)
	if errors.Cause(err) == sql.ErrNoRows{
		// 停止或者回滚
		return nil, errors.Wrap(err, "查询调用失败")
	}
	order, err := dao.New(goodsData.Id)

	if err != nil {
		 return nil, errors.Wrapf(err, "数据不存在")
	}

	return order, nil
}
