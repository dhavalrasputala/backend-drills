package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup" //always use go get or go install before importing package
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	exchanges := []string{"bitcoin", "ethurerum", "doge"}

	for _, exchange := range exchanges {
		ex := exchange //the exchange variable is changed every iteration and would eventually be doge so we capture it and every iteration is different

		g.Go(func() error { //g.GO is method provided by errorgroup similar to wg.add(1) it just does the same thing
			return fetchprice(ex, ctx)
		})
	}

	if err := g.Wait(); err != nil { //Mistake it should be "Wait" not "wait"
		fmt.Printf("Context Aborted :%v\n", err)
	} else {
		fmt.Print("All exchanges fetched sucessfully")
	}
}

func fetchprice(name string, ctx context.Context) error {
	fmt.Printf("[%s]Fetching price....\n", name)

	if name == "ethurerum" {
		fmt.Printf("[%s] API is down\n", name)
	}

	select {
	case <-time.After(3 * time.Second):
		fmt.Printf("[%s]Price Fetched Successfully", name)
		return nil
	case <-ctx.Done():
		fmt.Printf("[%s]Aborted Early :%v\n", name, ctx.Err())
		return ctx.Err()
	}
}
