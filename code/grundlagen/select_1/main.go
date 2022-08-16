package main

import "fmt"

func main() {
	c := make(chan string, 1)
	for i := 0; i < 10; i++ {
		select {
		case c <- "ping":
		case c <- "pong":
		case s := <-c:
			fmt.Println(s)
		default:
			fmt.Println("default")
		}
	}
}
