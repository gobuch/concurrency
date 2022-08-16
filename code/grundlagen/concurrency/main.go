package main

import (
	"fmt"
	"time"
)

// Was ist Concurrency?

// Quelle: https://go.dev/tour/concurrency/1

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}

func main() {
	go say("world")
	say("hello")
}
