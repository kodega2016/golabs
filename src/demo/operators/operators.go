package main

import "fmt"

func main() {
	a, b := 10, 20

	//arithmetic operator
	fmt.Println(a + b)
	fmt.Println(a - b)

	// increment / decrement operator
	a++
	fmt.Println(a)
	b--
	fmt.Println(b)

	// comparision operator
	fmt.Println(a > b)
	fmt.Println(a < b)

	// logical operator
	fmt.Println(a > 20 && b < 50)
	fmt.Println(a == b)
	fmt.Println(a != b)

}
