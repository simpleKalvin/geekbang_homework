package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/simpleKalvin/geekbang_homework/week01/errs"
)

type Goods struct {
	Id     int             `db:"Id"`
	Name   string          `db:"Name"`
	Amount uint8           `db:"amount"`
	Money  decimal.Decimal `db:"money"`
}

func GetOne(id int) (*Goods, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/shop"
	db,err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	goods := &Goods{}
	sqlString := "select * from goods where Id = ? limit 1"
	err = db.QueryRow(sqlString, id).Scan(&goods.Id, &goods.Name, &goods.Amount, &goods.Money)

	if errors.Is(err, sql.ErrNoRows) {
		errsData := new(errs.DataNotFound)
		return nil, errors.Wrapf(err, fmt.Sprintf("sqlString:%s, error: %v", sqlString, errsData))
	}

	return goods, nil
}
