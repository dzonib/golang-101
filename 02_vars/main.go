package main

import "fmt"

func main() {
	// creating variable using var
	// var name = "Dzoni"

	// int
	var age int32 = 100

	// bool
	var isCool = true
	isCool = false

	// shorthand
	// name := "King Kong,"
	// email := "kingkong@gmail.com"

	// shorter
	name, email := "King Kong", "kingkong@gmail.com"

	// float (64 by default)
	size := 1.23

	fmt.Println(name, email, age, isCool)
	fmt.Printf("%T\n", size)
}
