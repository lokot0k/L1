// демонстрация остновки горутины через нотифай из отдельного канала
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	stopChan := make(chan bool, 1) // буфер для гарантированного получения значения

	wg.Add(1)
	go func(wg *sync.WaitGroup, stopChan chan bool) { // типа отдельная функция
		defer wg.Done()
		for {
			select {
			case <-stopChan:
				fmt.Println("Горутина завершена")
				return
			default:
				fmt.Println("Горутина выполняется!")
				time.Sleep(200 * time.Millisecond) // имитация деятельности
			}
		}
	}(wg, stopChan)

	time.Sleep(5 * time.Second)
	stopChan <- true // сигнал остановки
	wg.Wait()
}
