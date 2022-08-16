package main

import (
	"fmt"
	"sync"
)

func greeter(wg *sync.WaitGroup, name string) {
	fmt.Println("Hallo ", name)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go greeter(wg, "routine 1")
	wg.Wait()
}
