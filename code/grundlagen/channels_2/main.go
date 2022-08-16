package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "eine Nachricht"
		ch <- "noch eine Nachricht"
	}()
	msg := <-ch
	fmt.Println(msg)
	msg = <-ch
	fmt.Println(msg)

}
