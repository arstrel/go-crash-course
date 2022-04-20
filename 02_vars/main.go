package main

import "fmt"

func main() {
	// using var
	var age int32 = 34
	const isCool = true

	// Shorthand only inside a function
	fullName := "Artem!"
	size := 1.2
	firstName, email := "Artem", "arstr@gmail.com"

	fmt.Println(fullName, firstName, email, age, isCool, size)
	fmt.Printf("%T\n", isCool)

}
