package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(s string) string {
	unicodeChars := []rune(s)
	for i := 0; i < len(unicodeChars)/2; i++ {
		j := len(unicodeChars) - i - 1
		unicodeChars[i], unicodeChars[j] = unicodeChars[j], unicodeChars[i]
	}
	return string(unicodeChars)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Строка для разворота: ")
	input, _ := reader.ReadString('\n')
	reversed := reverseString(input)

	fmt.Printf("Перевернутая строка: %s\n", reversed)
}
