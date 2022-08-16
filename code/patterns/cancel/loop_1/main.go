package main

import (
	"fmt"
	"time"
)

func main() {
	q := make(chan struct{})
	go do(q)
	time.Sleep(time.Second)
	close(q)
	fmt.Println("Ende")
}

func do(quit chan struct{}) {
	var counter int
	for {
		select {
		case <-quit:
			fmt.Println("quit function")
			return
		default:
			counter++
			fmt.Println("call nr: ", counter)
		}
	}
}
