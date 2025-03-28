package main

import "fmt"

func main() {
	// simple constant case with default
	s := "a"
	switch s {
	case "a":
		fmt.Println("it was an 'a'")
	default:
		fmt.Println("it wasn't an 'a'")
	}

	// if multiple conditions match only first match is picked up
	i := 5
	switch {
	case i > 0:
		fmt.Println("i is possitive")
	case i > 2:
		fmt.Println("i is greater than 2")
	case i > 5:
		fmt.Println("i is very big")
	default:
		fmt.Println("i is negative")
	}

	var a SomeInt
	a = A("hello")
	switch a.(type) {
	case A:
		fmt.Println("it's type A")
		a.Do()
	case B:
		fmt.Println("it's type B")
		a.Do()
	default:
		fmt.Println("unkown type")
	}
}

type A string
type B string

type SomeInt interface {
	Do()
}

func (A) Do() {
	fmt.Println("I'm doing as A")
}

func (B) Do() {
	fmt.Println("I'm doing as B")
}
