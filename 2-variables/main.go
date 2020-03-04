package main

import "fmt"

func main() {
	// Using var
	var name string = "Youssef"
	var age int = 20
	var isCool bool = true

	// We can set constants
	const Pi = 3.14
	// Pi = 3.2

	// Using the shorthand method
	gender := "male"
	email := "youssef@gmail.com"
	// gender, email := "male", "youssef@gmail.com"

	fmt.Println(name, age, isCool)
	fmt.Println(gender, email)
}
