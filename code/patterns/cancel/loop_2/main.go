package main

import "fmt"

func main() {
	done := make(chan struct{})
	go do(done)
	<-done
}

func do(done chan struct{}) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	close(done)
}
