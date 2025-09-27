package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите предложение: ")
	scanner.Scan()
	fmt.Printf("%s\n", reverseWords(scanner.Text()))
}

func reverseWords(s string) string {
	runes := []rune(s)
	start := 0
	reverseSubstringInPlace(runes, 0, len(runes)-1)
	for i, r := range runes {
		if unicode.IsSpace(r) {
			reverseSubstringInPlace(runes, start, i-1)
			start = i + 1
		}
	}
	reverseSubstringInPlace(runes, start, len(runes)-1)
	return string(runes)
}

func reverseSubstringInPlace(s []rune, start, end int) {
	for i := start; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}
