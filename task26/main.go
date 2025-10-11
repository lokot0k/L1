package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsCharsUnique(str string) bool {
	lower := strings.ToLower(str)
	used := make(map[rune]struct{})
	for _, char := range lower {
		if _, ok := used[char]; ok {
			return false
		}
		used[char] = struct{}{}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	if IsCharsUnique(input) {
		fmt.Println("В строке только уникальные символы")
	} else {
		fmt.Println("В строке есть повторяющиеся символы")
	}
}
