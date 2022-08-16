package main

import "fmt"

func main() {
	ch := gen()
	for nr := range ch {
		fmt.Println(nr)
	}
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
