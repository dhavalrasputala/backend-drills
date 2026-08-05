package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type StockPrice struct {
	Symbol string
	Price  float64
}

// fetchStockPrice simulates a slow network call to get a stock price.
func fetchStockPrice(symbol string) StockPrice {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	price := 100 + rand.Float64()*900
	return StockPrice{Symbol: symbol, Price: price}
}

// ---------- Approach 1: sync.WaitGroup ----------
// WaitGroup just blocks until N goroutines finish. It doesn't move data —
// each goroutine writes into its own index of a shared slice.
func fetchWithWaitGroup(symbols []string) []StockPrice {
	var wg sync.WaitGroup
	results := make([]StockPrice, len(symbols))

	for i, sym := range symbols {
		wg.Add(1)
		go func(i int, sym string) {
			defer wg.Done()
			results[i] = fetchStockPrice(sym) // safe: distinct index per goroutine
		}(i, sym)
	}

	wg.Wait() // block until all 3 goroutines call Done()
	return results
}

// ---------- Approach 2: buffered channel ----------
// The channel IS the communication: goroutines send results, main receives
// them. No shared slice, no manual synchronization needed.
func fetchWithChannel(symbols []string) []StockPrice {
	ch := make(chan StockPrice, len(symbols)) // buffered so sends don't block

	for _, sym := range symbols {
		go func(sym string) {
			ch <- fetchStockPrice(sym)
		}(sym)
	}

	results := make([]StockPrice, 0, len(symbols))
	for i := 0; i < len(symbols); i++ {
		results = append(results, <-ch) // receive as each result arrives
	}
	return results
}

func main() {
	symbols := []string{"AAPL", "GOOG", "TSLA"}

	fmt.Println("Using WaitGroup:")
	for _, s := range fetchWithWaitGroup(symbols) {
		fmt.Printf("  %s: $%.2f\n", s.Symbol, s.Price)
	}

	fmt.Println("\nUsing Channel:")
	for _, s := range fetchWithChannel(symbols) {
		fmt.Printf("  %s: $%.2f\n", s.Symbol, s.Price)
	}
}
