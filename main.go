package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		// How many concurrent activities are we creating
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) { // send-only channel
			if b, ok := queryCache(id, m); ok { // we only return values if we find in the cache
				ch <- b // Send message to the channel if we found the book in the cache
			}
			wg.Done()
		}(id, wg, m, cacheCh)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) { // send-only channel
			if b, ok := queryDatabase(id); ok { // always generates a result
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b // Send message to the channel if we found the book in the database
			}
			wg.Done()
		}(id, wg, m, dbCh)

		// create one goroutine per query to handle response
		go func(cacheCh, dbCh <-chan Book) { // receive-only channel
			select {
			case b := <-cacheCh: // receive message from cache
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh // because we don't want to block the second Goroutine
			case b := <-dbCh: // receive message from database
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)

		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond) // simulate database delay
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}
