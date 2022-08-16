package main

import (
	"fmt"
	"runtime"
	"time"
)

type job func(chan result)

type result struct {
	name string
	sum  int
	err  error
}

func startWorker(jobs chan job, results chan result) {
	go func() {
		fmt.Println("worker started")
		for j := range jobs {
			j(results)
		}
		fmt.Println("worker done")
	}()
}

func main() {
	jobs := make(chan job)
	results := make(chan result)
	go func() {
		for r := range results {
			fmt.Println(r)
		}
	}()
	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		startWorker(jobs, results)
	}
	job1 := func(resCh chan result) {
		r := result{
			name: "10+5",
			sum:  10 + 5,
			err:  nil,
		}
		resCh <- r
	}
	counter := 0
	job2 := func(resCh chan result) {
		counter++
		r := result{"Counter", counter, nil}
		resCh <- r
	}

	jobs <- job1
	jobs <- job2
	jobs <- job2
	jobs <- job2
	time.Sleep(time.Second)
}
