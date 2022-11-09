package main

import (
	"fmt"
)

func main() {
	role := "admin"

	switch role {
	case "admin":
		fmt.Println("admin")

	case "author":
		fmt.Println("author")

	default:
		fmt.Println("guest")
	}
}
