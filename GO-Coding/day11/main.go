// day11:-Task-Create an in-memory map representing an order book (Price -> Volume). Start 10 readers and 1 writer.
// Notice: Multiple readers can read simultaneously, but writes lock everything.(UNSOLVED)
package main

import (
	"math/rand"
)

type NewReader struct {
}

type stock struct {
	Price  int16
	Volume int16
}

func newstock(price, volume int16) *stock {
	return &stock{
		Price:  price,
		Volume: volume,
	}
}

func makestocks(num int) []*stock {
	stocks := make([]*stock, 0, num)
	for i := 0; i < num; i++ {
		vol := int16(rand.Intn(1000))
		price := int16(rand.Intn(10000))
		stocks = append(stocks, newstock(price, vol))
	}
	return stocks
}

func main() {
	stocs := makestocks(10)

}
