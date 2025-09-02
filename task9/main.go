package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	writerCount, multiplierCount := getInput()
	data := generateSlice(10000)
	inputChan := make(chan int, writerCount)
	processedChan := make(chan int, multiplierCount)
	go writeFromSlice(data, writerCount, inputChan)            // конкурретная запись в канал
	go outputResult(processedChan)                             // последовательный синхронный вывод, но в одном потоке
	processFromChan(inputChan, processedChan, multiplierCount) // конкурретный процессинг
}

func writeFromSlice(data []int, workersNum int, outputChan chan int) {
	wg := &sync.WaitGroup{}
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		// можем конкурретно обращаться к data, так как считаем его неизменяемым
		// если бы data кто-то писал/изменял, пришлось бы использовать RWMutex
		go func(index int) {
			defer wg.Done()
			for j := index; j < len(data); j += workersNum {
				outputChan <- data[j]
			}
		}(i)
	}
	wg.Wait()
	close(outputChan)
}

func processFromChan(inputChan chan int, outputChan chan int, workersNum int) {
	wg := &sync.WaitGroup{}
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range inputChan {
				outputChan <- num * 2
			}
		}()
	}
	wg.Wait()
	close(outputChan)
}

func outputResult(processedChan chan int) {
	for res := range processedChan {
		fmt.Println(res)
	}
}

func getInput() (int, int) {
	if len(os.Args) != 3 {
		panic("usage: ./task9 <writer_workers_num> <multiplier_workers_num>")
	}
	writerWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || writerWorkers < 1 {
		panic("invalid writer_workers_num, should be integer and > 0")
	}
	multiplierWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("invalid multiplier_workers_num, should be integer and > 1")
	}
	return writerWorkers, multiplierWorkers
}

func generateSlice(length int) []int {
	data := make([]int, length)
	for i := 0; i < length; i++ {
		data[i] = i
	}
	return data
}
