// демонстрация остановки горутины по условию
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	stopCondition := false

	wg.Add(1)
	// написали без захвата переменных из main, чтоб эмулировать самостоятельную функцию и посмотреть как она работает
	go func(wg *sync.WaitGroup, stopCondition *bool) {
		defer wg.Done()
		for !(*stopCondition) {
			fmt.Println("Горутина выполняется!")
			time.Sleep(200 * time.Millisecond) // эмуляция какаого-то процесса
		}
		fmt.Println("Горутина завершена")
	}(wg, &stopCondition)

	time.Sleep(2 * time.Second)
	stopCondition = true
	wg.Wait()
}
