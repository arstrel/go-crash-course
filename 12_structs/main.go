package main

import (
	"fmt"
	"strconv"
)

// Define struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Short sintax
type PersonShort struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " " + "and I am " + strconv.Itoa(p.age)
}

// hasBirthday methos (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	personShort1 := PersonShort{"Lee", "Bricks", "Miami", "M", 28}

	// initialize person using struct
	person1 := Person{firstName: "Artem", lastName: "Jumangy", city: "SF", gender: "M", age: 30}

	// Alternative
	person2 := Person{"Fiona", "Hutson", "SF", "F", 30}

	fmt.Println(personShort1)
	fmt.Println(person2.firstName)
	person1.age = 25
	fmt.Println(person1)

	fmt.Println(person1.greet())
	person2.hasBirthday()
	person2.hasBirthday()
	person2.getMarried("Stansfield")

	fmt.Println(person2.greet())

}
