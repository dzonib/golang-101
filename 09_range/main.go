package main

import "fmt"

func main() {
	// ids := []int{123,432,12,32,42,123}

	// loop them up
	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }

	// not using index
	// for _, id := range ids {
	// 	fmt.Printf("Index: %d\n", id)
	// }

	// Add all id's

	// sum:=0

	// for _, id := range ids {
	// 	sum += id
		// sum = sum + id
	// }
	
	// fmt.Println(sum)

	// range with maps
	emails := map[string]string{"Mr.king":"cooldude@king.com", "bob":"crappydude@gmail.com"}

	// for key, value := range emails {
	// 	fmt.Printf("%s: %s\n", key, value)
	// }

	// get keys (names)
	for k := range emails {
		fmt.Println("Name: " + k)
	}

	// get values (emails)
	for _, email := range emails {
		fmt.Println("Email: " +email)
	}
}
