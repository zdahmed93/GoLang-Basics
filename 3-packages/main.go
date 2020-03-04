package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/zdahmed93/go-in-hands/3-packages/utilities"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.Floor(7.9)) // 7
	fmt.Println(math.Ceil(7.9))  // 8
	fmt.Println(math.Sqrt(4))    // 2
	fmt.Println(strings.ToUpper("hi there"))
	fmt.Println(utilities.Uppercase("hi"))
}
