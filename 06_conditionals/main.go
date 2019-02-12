package main

import "fmt"

func main() {
	x := 10
	y := 10

	// if else
	if x <= y {
		fmt.Printf("%d is less or equal then %d\n", x, y)
	} else {
		fmt.Printf("%d is less then %d\n", y, x)
	}

	// else if
	color := "yellow"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("Color is not blue or red")
	}

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("Color is not blue or red")
	}
}
