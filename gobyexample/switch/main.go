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
}
