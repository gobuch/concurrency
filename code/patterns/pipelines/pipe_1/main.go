package main

import "fmt"

func main() {
	ch := gen()
	out := sq(ch)
	f := make(chan int)
	filter(out, f)
	for nr := range f {
		fmt.Println(nr)
	}
}

func sq(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func filter(in, out chan int) {
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
}

func gen() chan int {
	ch := make(chan int)
	go func() {
		i := 0
		for {
			ch <- i
			i++
			if i > 5 {
				close(ch)
				return
			}
		}
	}()
	return ch
}
