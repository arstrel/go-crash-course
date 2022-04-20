package main

import "fmt"

func main() {

	// Slice
	ids := []int{22, 35, 77, 89, 2, 112}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Loop through ids but without usage of index i
	for _, id := range ids {
		fmt.Printf(" ID: %d\n", id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum", sum)

	// Range with map

	data := map[string]string{"stan": "stan@gmail.com", "josh": "josh@gmail.com"}

	for k, v := range data {
		fmt.Printf(" %s: %s\n", k, v)
	}
}
