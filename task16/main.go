package main

import (
	"fmt"
)

func main() {
	arr := []int{8, 7, 6, 9, 10, -1, 4, 0, 7, 5, 11, 10, 3, -4, 9}
	fmt.Println("После сортировки:", quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left := make([]int, 0, len(arr)/2)
	mid := make([]int, 0, len(arr)/2)
	right := make([]int, 0, len(arr)/2)
	for _, elem := range arr {
		switch {
		case elem < pivot:
			left = append(left, elem)
		case elem == pivot:
			mid = append(mid, elem)
		case elem > pivot:
			right = append(right, elem)
		}
	}
	leftSorted := quickSort(left)
	rightSorted := quickSort(right)
	leftSorted = append(leftSorted, mid...)
	leftSorted = append(leftSorted, rightSorted...)
	return leftSorted
}
