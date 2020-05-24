package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch) // Receiving operation
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42 // Sending operation
		time.Sleep(5 * time.Millisecond)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
