package main

import "fmt"

func mathSwap(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func xorSwap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func assignmentSwap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func main() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Printf("\nВвели a=%d, b=%d\n", a, b)

	a, b = mathSwap(a, b)
	fmt.Printf("Поменяли местами вычитанием и сложением: a=%d, b=%d\n", a, b)
	a, b = xorSwap(a, b)
	fmt.Printf("Снова поменяли местами xor'ом: a=%d, b=%d\n", a, b)
	a, b = assignmentSwap(a, b)
	fmt.Printf("Снова поменяли местами присваиванием: a=%d, b=%d\n", a, b)
}
