package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		panic("Inappropriate number of arguments")
	}
	duration, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Duration should be an integer: %v", err))
	}
	inputChan := make(chan string)

	go handleInput(inputChan)
	populateInput(inputChan, time.After(time.Duration(duration)*time.Second))
}

func handleInput(inputChan chan string) {
	for val := range inputChan {
		fmt.Println(val)
	}
}

func populateInput(inputChan chan string, timer <-chan time.Time) {
	var data string
	for {
		select {
		case closeSignal := <-timer:
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
