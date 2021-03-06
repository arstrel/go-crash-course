package main

import "fmt"

func main() {
	x := 15
	y := 10

	// if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "blue"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not red or blue")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("switch: color is red")
	case "blue":
		fmt.Println("switch: color is blue")
	default:
		fmt.Println("switch: color is not red or blue")
	}

}
