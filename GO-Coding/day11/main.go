package main

import (
	"crypto/rand"
	"fmt"
	"sync"
)

type stock struct {
	Price int16
	Volume int16
}

func newstock(price,volume int16) *stock {
	return &stock{
		Price: price,
		Volume: volume,
	}
}

func (s stock) makestocks(num int) {
	for i=0;i<=num;i++{
		newstock(s.Price:)
	}
}

func main() {

}
