package main

import "fmt"

func main() {

	a := 5
	b := &a

	fmt.Println(a, b)
	// b is printed as memory address => 0xc000050058

	// check type
	fmt.Printf("%T\n", a)
	// a is int
	fmt.Printf("%T\n", b)
	// b is *int -> int and *int are different types
	// one with star represent pointer

	// USE * TO READ VAL FROM ADDRESS OR IT RETURN MEMORY ADDRESS
	fmt.Printf("%d\n", *b)

	// change val with a pointer
	*b = 10

	fmt.Printf("%d\n", a)
}
