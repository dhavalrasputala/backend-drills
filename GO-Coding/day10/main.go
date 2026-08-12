package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	ctxtimeout, cancel := context.WithTimeout(context.Background(), 2*time.Second) //Change the timeout value to see the result variable
	defer cancel()
	go dosomething(ctxtimeout, ch)

	select {
	case <-ctxtimeout.Done():
		fmt.Printf("Context Cancelled :%v", ctxtimeout.Err())
	case result := <-ch:
		fmt.Printf("Recived Value :%v", result)
	}
}

func dosomething(ctx context.Context, channel chan string) {
	fmt.Print("Sleeping")
	time.Sleep(5 * time.Second)
	fmt.Print("waking up!")
	channel <- "did something here"
}
