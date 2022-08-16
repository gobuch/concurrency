package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	c := make(chan string)
	wg.Add(2)
	go sender(wg, c)
	go receiver(wg, c)
	wg.Wait()
}

func sender(wg *sync.WaitGroup, c chan string) {
	c <- "Hallo receiver"
	wg.Done()
}

func receiver(wg *sync.WaitGroup, c chan string) {
	msg := <-c
	fmt.Println("r:", msg)
	wg.Done()
}
