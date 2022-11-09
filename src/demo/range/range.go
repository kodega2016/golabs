package main

import "fmt"

func main() {
	users := []string{
		"Khadga Shrestha", "Nishuka Shrestha", "Menuka Shrestha",
	}

	for i, name := range users {
		fmt.Println(i, name)
	}
}
