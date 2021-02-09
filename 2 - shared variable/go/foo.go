// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"runtime"
)

func number_server(add <-chan int, sub <-chan int, read chan<- int) {
	var number = 0

	// This for-select pattern is one you will become familiar with...
	for {
		select {
		case <-add:
			number++
		case <-sub:
			number--
		case read <- number:
			read <- number
		}
	}
}

func incrementer(add chan<- int, finishedIncr chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add <- 1
	}
	finishedIncr <- true
}

func decrementer(sub chan<- int, finishedDecr chan<- bool) {
	for j := 0; j < 1000000+1; j++ {
		sub <- 1
	}
	finishedDecr <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	add := make(chan int)
	sub := make(chan int)
	finishedIncr := make(chan bool)
	finishedDecr := make(chan bool)
	read := make(chan int)

	go number_server(add, sub, read)
	go incrementer(add, finishedIncr)
	go decrementer(sub, finishedDecr)
	<-finishedIncr
	<-finishedDecr

	fmt.Println("The magic number is:", <-read)
}
