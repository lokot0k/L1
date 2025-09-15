package main

import (
	"fmt"
)

func createValue(choice int) interface{} {
	switch choice {
	case 1:
		return 0

	case 2:
		return "agsagd"

	case 3:
		return false

	case 4:
		return make(chan int)

	default:
		return nil
	}
}

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Тип: int")
	case string:
		fmt.Println("Тип: string")
	case bool:
		fmt.Println("Тип: bool")
	case chan int:
		fmt.Println("Тип: chan int")
	default:
		fmt.Println("Неподдерживаемый тип")
	}
}

func main() {
	var choice int
	fmt.Print("Введите число 1 - 4:\n1-int\n2-string\n3-bool\n4-chan\n")
	_, err := fmt.Scan(&choice)
	if err != nil {
		panic(err)
	}
	value := createValue(choice)
	printType(value)
}
