package main

import "fmt"

func main() {

	// Define map
	emails := make(map[string]string)

	// Assign key-values
	emails["bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["bob"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "bob")
	fmt.Println(emails)

	// declare map and add key-values

	data := map[string]string{"stan": "stan@gmail.com", "josh": "josh@gmail.com"}
	fmt.Println(data)

}
