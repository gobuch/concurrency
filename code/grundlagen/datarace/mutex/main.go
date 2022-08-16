package main

import (
	"fmt"
	"sync"
)

// Diesen Code mit dem Race Detector ausführen!
// go run -race main.go

func main() {
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)
	mutex := &sync.Mutex{}
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
