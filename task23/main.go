package main

import "fmt"

func removeFromSlice(slice []int, ind int) []int {
	return append(slice[:ind], slice[ind+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	reducedSlice := removeFromSlice(slice, 3)
	fmt.Println(reducedSlice)
}
