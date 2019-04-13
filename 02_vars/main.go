package main

import "fmt"

// can't declare outside of a function
// kingkong := true

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Kingkong"

	// short declaration
	name := "Kingkong"

	var age int32 = 100

	// constant
	const isCool = false

	size := 225.883

	// declare more then one variable
	username, password := "god", "123456"

	fmt.Println(name, age, username, password)

	// https://golang.org/pkg/fmt/
	// find type of given value
	fmt.Printf("%T\n", size)
}
