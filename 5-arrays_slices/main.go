package main

import "fmt"

func main() {
	// Arrays
	// Declare
	var colorsArray [3]string

	// Assign values
	colorsArray[0] = "red"
	colorsArray[1] = "green"
	colorsArray[2] = "blue"

	// Declare and assign
	evenNumbers := []int{2, 4, 6, 8}

	// if we wanna add a new item
	evenNumbers = append(evenNumbers, 10)

	fmt.Println(colorsArray)
	fmt.Println(evenNumbers)

	// Slices
	colorsSlice := []string{"pink", "grey", "purple", "yellow"}
	fmt.Println(colorsSlice)
	fmt.Println(len(colorsSlice))
	fmt.Println(colorsSlice[0:2])
}
