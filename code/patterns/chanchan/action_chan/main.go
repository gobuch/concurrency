package main

import "fmt"

type task struct {
	a, b   int
	result chan int
	err    chan error
}

func main() {
	ch := make(chan task)
	go add(ch)
	go sub(ch)
	resCh := make(chan int)
	ch <- task{a: 3, b: 4, result: resCh, err: make(chan error)}
	ch <- task{a: 3, b: 4, result: resCh, err: make(chan error)}
	ch <- task{a: 3, b: 4, result: resCh, err: make(chan error)}
	ch <- task{a: 3, b: 4, result: resCh, err: make(chan error)}
	close(ch)
	fmt.Println(<-resCh)
	fmt.Println(<-resCh)
	fmt.Println(<-resCh)
	fmt.Println(<-resCh)
}

func add(tasks chan task) {
	for task := range tasks {
		r := task.a + task.b
		task.result <- r
	}
}

func sub(tasks chan task) {
	for task := range tasks {
		fmt.Println("sub")
		r := task.a - task.b
		task.result <- r
	}
}
