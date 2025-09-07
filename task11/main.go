package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите элементы первого множества (через пробел):")
	set1 := getInputSet()
	fmt.Println("Введите элементы второго множества (через пробел):")
	set2 := getInputSet()
	intersection := getIntersection(set1, set2)
	fmt.Println("Пересечение множеств:")
	for elem := range intersection {
		fmt.Printf("%d ", elem)
	}
	fmt.Println()
}

func getInputSet() map[int]struct{} {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	elems := strings.Split(input, " ")
	set := make(map[int]struct{})
	for _, elem := range elems {
		num, err := strconv.Atoi(elem)
		if err != nil {
			panic("необходимо вводить только целочисленные числа")
		}
		set[num] = struct{}{}
	}
	return set
}

func getIntersection(set1, set2 map[int]struct{}) map[int]struct{} {
	intersect := make(map[int]struct{})
	for elem := range set1 {
		if _, exists := set2[elem]; exists {
			intersect[elem] = struct{}{}
		}
	}
	return intersect
}
