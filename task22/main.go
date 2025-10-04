package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

type BigInt struct {
	value *big.Int
}

func NewBigInt(num string) *BigInt {
	b := &BigInt{value: new(big.Int)}
	b.value.SetString(num, 10)
	return b
}

func (b *BigInt) Add(other *BigInt) *BigInt {
	res := new(big.Int).Add(b.value, other.value)
	return &BigInt{value: res}
}

func (b *BigInt) Subtract(other *BigInt) *BigInt {
	res := new(big.Int).Sub(b.value, other.value)
	return &BigInt{value: res}
}

func (b *BigInt) Multiply(other *BigInt) *BigInt {
	res := new(big.Int).Mul(b.value, other.value)
	return &BigInt{value: res}
}

func (b *BigInt) Divide(other *BigInt) (*BigInt, error) {
	if other.value.Sign() == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	res := new(big.Int).Div(b.value, other.value)
	return &BigInt{value: res}, nil
}

func (b *BigInt) String() string {
	return b.value.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("a: ")
	scanner.Scan()
	a := NewBigInt(strings.TrimSpace(scanner.Text()))
	fmt.Print("b: ")
	scanner.Scan()
	b := NewBigInt(strings.TrimSpace(scanner.Text()))
	if a == nil || (a.value.Sign() == 0 && strings.TrimSpace(scanner.Text()) != "0") {
		fmt.Println("Invalid input a")
		return
	}
	if b == nil || (b.value.Sign() == 0 && strings.TrimSpace(scanner.Text()) != "0") {
		fmt.Println("Invalid input b")
		return
	}

	sum := a.Add(b)
	fmt.Printf("Sum: %s\n", sum)

	difference := a.Subtract(b)
	fmt.Printf("Sub: %s\n", difference)

	product := a.Multiply(b)
	fmt.Printf("Mul: %s\n", product)

	div, err := a.Divide(b)
	if err != nil {
		fmt.Printf("Div error: %s\n", err)
	} else {
		fmt.Printf("Div: %s\n", div)
	}
}
