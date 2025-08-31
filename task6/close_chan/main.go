// демонстрация остановки горутины по закрытию канала
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	inputChan := make(chan int)

	wg.Add(1)
	go func(wg *sync.WaitGroup, inputChan <-chan int) { // эмулируем самостоятельную функцию, поэтому передаем параметры, а не захватываем
		defer wg.Done()
		for data := range inputChan {
			fmt.Println(data)
		}
		fmt.Println("Горутина завершена")
	}(wg, inputChan)
	timer := time.After(5 * time.Second)
	i := 0
	for {
		select {
		case <-timer:
			close(inputChan)
			wg.Wait()
			return

		default:
			time.Sleep(200 * time.Millisecond)
			inputChan <- i
			i++
		}

	}

}
