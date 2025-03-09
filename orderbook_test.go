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

	buyOrder1 := NewOrder(true, 10)
	buyOrder2 := NewOrder(true, 2000)
	ob.PlaceOrder(18_000, buyOrder1)
	ob.PlaceOrder(19_000, buyOrder2)

	for i := 0; i < len(ob.Bids); i++ {
		fmt.Println(ob.Bids[i])
	}
}
