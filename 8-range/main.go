package main

import "fmt"

func main() {
	colors := []string{"red", "green", "blue"}

	// Loop through the slice of colors
	for index, item := range colors {
		fmt.Println(index, item)
	}

	// Not using index
	colorsHexaCode := []string{"#FF0000", "#00FF00", "#0000FF"}
	for _, item := range colorsHexaCode {
		fmt.Println(item)
	}
}
