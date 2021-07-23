package order

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	order := New(1)
	fmt.Println(order)
}
