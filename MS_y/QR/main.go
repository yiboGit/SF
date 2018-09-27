package main

import (
	"fmt"
)

type animalCategory struct {
	category     string
	categoryName string
}

type animal struct {
	name string
	animalCategory
}

func main() {
	var a = animal{
		name: "cat",
	}
	b := a.setName("dog")
	fmt.Println(b)
}

func (a animal) setName(n string) animal {
	a.name = n
	return a
}
