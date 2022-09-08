package main

import (
	"fmt"
)

type demo struct {
	data int
}

func demoFunction(data int) *demo {
	d := demo{data: data}
	return &d
}

func insufficientStack() {
	s := make([]int, 10000, 10000)
	for index := range s {
		s[index] = index
	}
}

func main() {
	demo := demoFunction(11)
	println(demo)

	insufficientStack()

	fmt.Println("dynamic type escape")

	BuildLibrary()
}

var AvailableForLoan []*Book

type Book struct {
	Mice           int
	Men            int
	Classification string
}

func BuildLibrary() {
	bNoEscape := Book{Mice: 99, Men: 3}
	AddToCollectionNoEscape(&bNoEscape)

	bEscape := Book{Mice: 99, Men: 3}
	AddToCollectionEscape(&bEscape)
}

func AddToCollectionNoEscape(b *Book) {
	b.Classification = "fiction"
}

func AddToCollectionEscape(b *Book) {
	AvailableForLoan = append(AvailableForLoan, b)
}
