package dao

import (
	"fmt"
	_ "github.com/pkg/errors"
	"testing"
)

func TestGoodsGetOne(t *testing.T) {
	fmt.Println(123)
	_, err := GetOne(1)

	if err != nil {
		fmt.Println(err)
	}
}
