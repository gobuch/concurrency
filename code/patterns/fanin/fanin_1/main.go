package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := gen()
	c2 := gen()
	c3 := gen()
	m := merge(c1, c2, c3)
	for n := range m {
		fmt.Println(n)
	}
}

func merge(cs ...chan int) chan int {
	wg := &sync.WaitGroup{}
	out := make(chan int)
	output := func(c chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen() chan int {
	ch := make(chan int)
	go func(c chan int) {
		for i := 1; i <= 3; i++ {
			c <- i
		}
		close(c)
	}(ch)
	return ch
}
