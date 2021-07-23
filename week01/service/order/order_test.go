package order

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	order, err := New(2)
	fmt.Println(err)
	fmt.Println(order)
}
