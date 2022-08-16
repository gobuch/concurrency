package main

import (
	"fmt"
	"time"
)

func main() {
	limit := 5
	sem := make(chan bool, limit)
	hugeSlice := make([]bool, 10)
	for i, task := range hugeSlice {
		sem <- true
		go func(task bool, nr int) {
			fmt.Println("Working: ", nr)
			time.Sleep(time.Second * time.Duration(nr))
			fmt.Println("Ready: ", nr)
			<-sem
		}(task, i)
	}
}
