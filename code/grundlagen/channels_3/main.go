package main

import "fmt"

func main() {
	s := add(1, 2)
	fmt.Println(s)
	in := make(chan int)
	out := make(chan int)
	go wrapper(in, out)
	in <- 2
	in <- 4
	s = <-out
	fmt.Println(s)
}

func wrapper(in chan int, out chan int) {
	a := <-in
	b := <-in
	out <- add(a, b)
}

func add(a, b int) int {
	return a + b
}
