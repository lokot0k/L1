// демонстрация остановки горутин через goexit
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			fmt.Println("Горутина работает")
			time.Sleep(200 * time.Millisecond)
			if i == 3 {
				defer fmt.Println("Горутина завершена")
				runtime.Goexit() // завершаем горутину, но ничего больше
			}
		}
	}(wg)

	wg.Wait()
}
