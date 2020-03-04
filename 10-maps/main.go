package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	firstMap := make(map[string]string)
	firstMap["red"] = "#FF0000"
	firstMap["green"] = "#00FF00"
	fmt.Println(firstMap)
}
