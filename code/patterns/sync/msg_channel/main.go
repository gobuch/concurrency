package main

import (
	"fmt"
)

func greeter(msg chan string, name string) {
	msg <- fmt.Sprintf("Hallo %s\n", name)
}

func main() {
	msgchan := make(chan string)
	go greeter(msgchan, "Alice")
	fmt.Println(<-msgchan)
}
