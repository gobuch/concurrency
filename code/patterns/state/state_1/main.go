package main

import "fmt"

type kontoMachine struct {
	saldo int
	inCh  chan int
	outCh chan int
	quit  chan struct{}
}

func (km *kontoMachine) run() {
	go func() {
		fmt.Println("km started...")
		for {
			select {
			case in := <-km.inCh:
				km.saldo = km.saldo + in
			case out := <-km.outCh:
				km.saldo = km.saldo - out
			case <-km.quit:
				fmt.Println("km ended")
				return
			}
			fmt.Printf("Saldo: %d\n", km.saldo)
		}
	}()
}

func newKontoMachine() *kontoMachine {
	return &kontoMachine{
		saldo: 0,
		inCh:  make(chan int),
		outCh: make(chan int),
		quit:  make(chan struct{}),
	}
}

func main() {
	km := newKontoMachine()
	km.run()
	km.inCh <- 10
	km.inCh <- 10
	km.inCh <- 10
	km.outCh <- 20
	km.quit <- struct{}{}
}
