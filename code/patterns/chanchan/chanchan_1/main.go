package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go routine1(ch)
	fmt.Println(<-ch)
}

func routine1(result chan string) {
	tasks := make(chan chan string)
	go routine2(tasks)
	task := make(chan string)
	tasks <- task
	result <- <-task
}

func routine2(tasks chan chan string) {
	for resCh := range tasks {
		// ...
		time.Sleep(time.Second)
		resCh <- "Ergebnis routine 2"
	}
}
