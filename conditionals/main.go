package main

import "fmt"

func main() {
	if true {
		fmt.Println("Always True")

	}

	b := 4

	// a is short scoped variable
	if a := b; a == 3 {
		fmt.Println("a is 3")
	} else if a == 2 {
		fmt.Println("a is 2")
	} else if a == 4 {
		fmt.Println("a is 4")
	} else {
		fmt.Println("Default")
	}

	// a is short scoped variable
	switch a := b; a {
	case 3:
		fmt.Println("a is 3")
	case 2:
		fmt.Println("a is 2")
	case 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("Default")
		fmt.Println("Value")
	}
}
