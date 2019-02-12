package main

import "fmt"

func main() {
	// define map
	// emails := make(map[string]string)

	// assign key values
	// emails["bob"] = "bob@gmail.com"
	// emails["yoy"] = "das@gmail.com"
	// emails["Mr.king"] = "coolbro@gmail.com"

	// Declare map and add key-values

	emails := map[string]string{"Mr.king":"cooldude@king.com", "bob":"crappydude@gmail.com"}

	// add more
	emails["yoy"] = "das@gmail.com"

	fmt.Println(len(emails))
	fmt.Println(emails["Mr.king"])

	// delete from map
	delete(emails, "bob")
	fmt.Println(emails)

}
