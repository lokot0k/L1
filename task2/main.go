package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	for _, i := range input {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Printf("Power of 2 of number %d is: %d\n", num, num*num)
		}(i)
	}
	wg.Wait()
}
