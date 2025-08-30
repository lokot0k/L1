package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		panic("Inappropriate number of arguments")
	}
	workersNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Workers quanity should be an integer: %v", err))
	}
	inputChan := make(chan string, workersNum)
	sigChan := make(chan os.Signal, 1) // канал для получения ctrl+c, буферизируем, чтобы избежать блокирующих вызовов
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	wg := &sync.WaitGroup{} // wg для обработчиков вывода, для того, чтоб не пропали уже положенные в канал данные
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go handleInput(inputChan, wg)
	}
	populateInput(inputChan, sigChan)
	wg.Wait()
}

func populateInput(inputChan chan string, sigChan chan os.Signal) {
	var data string
	for {
		select {
		case closeSignal := <-sigChan:
			fmt.Println("Получили сигнал на выход: ", closeSignal)
			close(inputChan) // закрываем канал, прекращаем отправку данных
			return
		default:
			err := faker.FakeData(&data)
			if err != nil {
				data = err.Error()
			}
			inputChan <- data
		}
	}
}

func handleInput(inputChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range inputChan { // inputChan - буферизированный, значит после sigInt дообработаем то, что в нем лежит
		fmt.Println(val)
	}
}
