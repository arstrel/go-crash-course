package main

import "fmt"

func main() {
	// Arrays have to be of a fixed length
	var fruit [2]string

	// Assign values
	fruit[0] = "Apple"
	fruit[1] = "Orange"

	// Declare and assign values
	fruitArr := [2]string{"Banana", "Peach"}

	// Slices
	fruitSlice := []string{"Banana ", "Peach ", "Grape", "Cherry"}

	fmt.Println(fruitArr, fruitSlice, len(fruitSlice))
	// starts at 1 and stops at 2 (includes item at index 2)
	// Output: [Peach Grape]
	fmt.Println(fruitSlice[1:3])
}
