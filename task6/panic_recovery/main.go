// демонстрация остановки горутины через панику
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Восстановление после:", r)
			}
		}()

		for i := 0; i < 1000; i++ {
			fmt.Println("Горутина работает")
			time.Sleep(200 * time.Millisecond)
			if i == 3 {
				panic("Паника!")
			}
		}
	}()

	wg.Wait()
}
