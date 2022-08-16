package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 5; i++ {
		go zuhörer(ch)
		time.Sleep(time.Second)
	}
	ch <- "eine Nachricht"
	ch <- "zweite Nachricht"
	close(ch)
	time.Sleep(time.Second)
	fmt.Println("Ende")
}

func zuhörer(ch chan string) {
	fmt.Println("warte...")
	msg := <-ch
	fmt.Println("msg: ", msg)
	fmt.Println("goroutine Ende")
}
