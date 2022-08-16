package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})
	go func() {
		// doSomething()
		close(done)
	}()
	fmt.Println("warte")
	<-done
	fmt.Println("fertig")
}
