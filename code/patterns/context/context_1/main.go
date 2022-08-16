package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	nrChan := gen(ctx)
	for n := range nrChan {
		fmt.Println(n)
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Beenden der Goroutine")
				close(dst)
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
