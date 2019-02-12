package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// slice
	fruitSlice := []string{"Apple", "Orange", "Grape", "Chearry"}

	fmt.Println(len(fruitSlice))

	fmt.Println(fruitSlice[1:3])
}
