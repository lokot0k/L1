package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ConcurrentCounter atomic.Int64

func (c *ConcurrentCounter) Inc() {
	(*atomic.Int64)(c).Add(1)
}

func (c *ConcurrentCounter) Value() int64 {
	return (*atomic.Int64)(c).Load()
}

var numberOfIncrementers int = 100000

func main() {
	counter := &ConcurrentCounter{}
	var wg sync.WaitGroup
	for i := 0; i < numberOfIncrementers; i++ {
		wg.Add(1)
		go func(counter *ConcurrentCounter) {
			defer wg.Done()
			counter.Inc()
		}(counter)
	}
	wg.Wait()
	fmt.Println(counter.Value())
}
