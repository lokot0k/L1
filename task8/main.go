package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, index, err := getInput()
	if err != nil {
		panic(err)
	}
	output := switchBit(num, index)
	fmt.Printf("Result after switchng %d'th bit in %d: %d\n", index, num, output)
}

func getInput() (int64, int, error) {
	var err error
	if len(os.Args) != 3 {
		panic("Usage: ./task8 <num> <index>\n <index> should be from 0 to 63 inclusive")
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("num should be an integer: %v", err))
	}
	index, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(fmt.Sprintf("index should be an integer: %v", err))
	}
	if index < 0 || index >= 64 {
		err = fmt.Errorf("index should be between 0 and 63 inclusive: %v", err)
	}
	return int64(num), index, err
}

func switchBit(num int64, i int) int64 {
	return num ^ (1 << i)
}
