package main

import (
	"fmt"
	"time"
)

func producer(boundedBuffer chan<- int) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		boundedBuffer <- i
	}

}

func consumer(boundedBuffer <-chan int) {

	time.Sleep(1 * time.Second)
	for {
		i := 0 //TODO: get real value from buffer
		i = <-boundedBuffer
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}

}

func main() {

	boundedBuffer := make(chan int, 5)

	go consumer(boundedBuffer)
	go producer(boundedBuffer)

	select {}
}
