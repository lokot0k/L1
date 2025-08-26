package main

import "fmt"

type Human struct {
	Name    string
	Surname string
}

func (h Human) PrintSomething() {
	fmt.Printf("Human %s %s is something!\n", h.Name, h.Surname)
}

func (h Human) PrintSurname() {
	fmt.Printf("%s\n", h.Surname)
}

func (h Human) PrintName() {
	fmt.Printf("%s\n", h.Name)
}

type Action struct {
	Human
}

func main() {
	h := Human{"Vasiliy", "Vasiliev"}
	a := Action{h}
	a.PrintName()
	a.PrintSurname()
	a.PrintSomething()
}
