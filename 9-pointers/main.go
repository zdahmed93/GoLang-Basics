package main

import "fmt"

func main() {
	// a := 5
	// b := &a

	// fmt.Println(a, b)
	// fmt.Printf("%T\n", b)

	// //  Use * to read val from address
	// fmt.Println(*b)
	// fmt.Println(*&a)

	// // Change val with pointer
	// *b = 10
	// fmt.Println(a)

	a := 3
	addOne(3)
	fmt.Println("a", a)

	addTen(&a)
	fmt.Println("a", a)

}

func addOne(i int) {
	i = i + 1
}

func addTen(p *int) {
	*p = *p + 10
}
