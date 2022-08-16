package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan []byte, 1)
	go func() {
		result <- networkRequest()
	}()
	select {
	case data := <-result:
		fmt.Println(data)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

func networkRequest() []byte {
	fmt.Println("request started...")
	time.Sleep(2 * time.Second)
	return []byte{1, 2}
}
