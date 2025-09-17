package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
	return strings.Repeat("0", size)
}

func someFunc() {
	v := createHugeString(1 << 10)
	buffer := make([]byte, 100)
	copy(buffer, v)
	justString = string(buffer)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
