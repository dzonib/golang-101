package main

import "fmt"

func main() {

	// i := 1

	// Long method

	// for i <= 10 {
	// 	fmt.Println(i)
	// 	// i = i + 1
	// 	i++
	// }

	// Short method

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// fizzbuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("Fizbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
