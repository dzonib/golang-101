package main

import (
	"fmt"
	"strconv"
)

//Person define struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// if you dont change anything it only recives values so its value reciver
// Greeting method (value reciver)
// p is like this in js, greet is name, string is return value
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " +
		strconv.Itoa(p.age) + " years old."
}

// if you change value its pointer reciver
// hasBirthay method (pointer reciver)
func (p *Person) hasBirthay() {
	p.age++
}

// gets married reciver
// change last name if she gets merried
func (p *Person) gotMerried(luckyDude Person) {
	if p.gender == "female" {
		p.lastName = luckyDude.lastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "King", lastName: "Kong", city: "Jungle", gender: "male", age: 5283}
	person3 := Person{firstName: "Yoyo", lastName: "Animal", city: "Jungle", gender: "male", age: 53}

	// alternative
	person2 := Person{"Queen", "Schmong", "Jungle", "female", 5283}

	person1.age = person1.age + 100
	person1.age++
	fmt.Println(person1)
	person2.hasBirthay()
	person2.gotMerried(person3)
	fmt.Println(person2.greet())
}
