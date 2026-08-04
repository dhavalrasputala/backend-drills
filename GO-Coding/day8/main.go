// Day 8:-Task:Goroutine Leaks: Write a function that spawns a goroutine blocking forever on an unbuffered channel.
// Call it 100 times. Print runtime.NumGoroutine().
// Notice: How orphaned goroutines silently eat memory.
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan struct{})
	for i := 0; i <= 100; i++ {
		go worker(i, ch)
	}
	//for i := 0; i <= 150; i++ { //Mistake :- TO cause a go routine leak make sure the main function never reads the channel
	//<-ch
	//}
	time.Sleep(2 * time.Second) //Mistake:- give some time for Goroutines to exists
	remaining := runtime.NumGoroutine()
	fmt.Printf("remaining Go routines are:- %v", remaining)
}

func worker(id int, ch chan<- struct{}) {
	fmt.Printf("Worker %d started", id)
	ch <- struct{}{}
}
