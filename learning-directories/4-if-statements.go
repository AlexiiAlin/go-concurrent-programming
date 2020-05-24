package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		if msg, isOpened := <-ch; isOpened {
			fmt.Println(msg, isOpened)
		}
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 0
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
