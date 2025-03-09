package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)

	buyOrder1 := NewOrder(true, 5)
	buyOrder2 := NewOrder(true, 8)
	buyOrder3 := NewOrder(true, 10)
	l.AddOrder(buyOrder1)
	l.AddOrder(buyOrder2)
	l.AddOrder(buyOrder3)

	l.DeleteOrder(buyOrder2)

	fmt.Println(l)
}

func TestOrderbook(t *testing.T) {
	ob := NewOrderbook()
}
