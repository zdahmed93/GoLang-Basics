package main

import "fmt"

func main() {
	fmt.Println(greeting("ahmed"))
	fmt.Println(getSum(2, 5))
}

func greeting(s string) string {
	return "Hello " + s
}

func getSum(a, b int) int {
	return a + b
}
