package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("wrote", i, "to channel")
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go producer(ch)
	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("reading", v, "from channel")
		time.Sleep(2 * time.Second)
	}
}
