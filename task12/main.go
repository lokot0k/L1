package main

import "fmt"

type Set struct {
	elements map[string]interface{}
}

func NewSet() *Set {
	return &Set{
		elements: make(map[string]interface{}),
	}
}

func (s *Set) Add(element string) {
	s.elements[element] = nil
}

func (s *Set) GetElements() []string {
	keys := make([]string, 0)
	for k := range s.elements {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSet()
	for _, word := range input {
		set.Add(word)
	}
	fmt.Printf("%v\n", set.GetElements())
}
