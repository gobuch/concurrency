package main

import (
	"fmt"
	"sync"
)

func greeter(wg *sync.WaitGroup, name string) {
	fmt.Printf("Hallo %s\n", name)
	wg.Done()
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go greeter(wg, "Alice")
	wg.Wait()
}
