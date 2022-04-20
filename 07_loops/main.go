package main

import "fmt"

func main() {

	// Long method
	i := 1
	for i <= 5 {
		fmt.Println(i)
		// i = i + 1
		i++
	}

	// Short method
	for i := 1; i <= 7; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz
	for i := 1; i < 31; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
