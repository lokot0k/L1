package main

import (
	"bytes"
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
	return strings.Repeat("0", size)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(bytes.Clone([]byte(v[:100])))
}

func main() {
	someFunc()
	fmt.Println(justString)
}
