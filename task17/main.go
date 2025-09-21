package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{8, 7, 6, -1, 4, 0, 5, 11, 10, 3, -4, 9}
	sort.Ints(arr) // сортировка необходима для бин поиска
	fmt.Println("После сортировки: ", arr)
	elemToFind := 5
	fmt.Println("Индекс запрашиваемого эелмента ", elemToFind, ": ", binarySearch(arr, elemToFind))
}

func binarySearch(arr []int, elem int) int {
	leftBound := 0
	rightBound := len(arr) - 1
	for leftBound <= rightBound {
		mid := leftBound + (rightBound-leftBound)/2
		if arr[mid] == elem {
			return mid
		}
		if arr[mid] > elem {
			rightBound = mid - 1
		} else {
			leftBound = mid + 1
		}
	}

	return -1
}
