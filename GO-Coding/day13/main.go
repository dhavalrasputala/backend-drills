// Day13:-Task:Worker Pool: Create a pool of 3 workers. Send 10,000 end-of-settlement jobs to a channel.
// Notice: Bounded concurrency prevents overwhelming the settlement database.
package main

import (
	"fmt"
	"sync"
)

type job struct {
	Id   int
	Data string
}

func main() {
	jobs := make(chan job, 10000)
	results := make(chan string, 10000)
	var wg sync.WaitGroup //Mistake:- wrote var wg *sync.WaitGroup this creates a null pointer not an actual waitgroup use var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for i := 1; i <= 10000; i++ { //here we are sending jobs to the channel we can do it in worker function also
		jobs <- job{Id: i, Data: fmt.Sprintf("task %d", i)}
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()
	for res := range results {
		fmt.Println(res)
	}
}

func worker(id int, jobs <-chan job, result chan<- string, wg *sync.WaitGroup) { //Mistake:- jobs chan<-job means send only jobs <-chan job means recive only jobs chan job means send & recieve both
	defer wg.Done()
	for i := range jobs {
		result <- fmt.Sprintf("worker %v Completing %s", id, i.Data)
	}
}
