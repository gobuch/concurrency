package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	name := "Alice"
	wg.Add(1)
	go func(wg *sync.WaitGroup, name string) {
		fmt.Println("Hallo", name)
		wg.Done()
	}(wg, name)
	name = "Bob"
	wg.Wait()
}
