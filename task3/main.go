package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"os"
	"strconv"
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
	for i := 0; i < workersNum; i++ {
		go handleInput(inputChan)
	}
	populateInput(inputChan)
}

func populateInput(inputChan chan string) {
	var data string
	for {
		err := faker.FakeData(&data)
		if err != nil {
			data = err.Error()
		}
		inputChan <- data
	}
}

func handleInput(inputChan chan string) {
	for {
		fmt.Println(<-inputChan)
	}
}
