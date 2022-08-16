package main

import "fmt"

func main() {
	c := make(chan int)
	go foo(c)
	fmt.Println(<-c)

}

func foo(ch chan<- int) {
	// ...
	ch <- 1
	// ...
}
