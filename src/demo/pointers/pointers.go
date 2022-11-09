package main

import "fmt"

func main() {
	var p *int
	a := 10
	p = &a

	fmt.Println("Address:", p)
	fmt.Printf("Value with %d address is %d", p, *p)
}
