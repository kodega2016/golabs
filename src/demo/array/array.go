package main

import "fmt"

func main() {
	users := []string{
		"Khadga Shrestha", "Nishuka Shrestha", "Menuka Shrestha",
	}

	fmt.Println(users)

	fmt.Println("total user:", len(users))
	fmt.Println("first user:", users[0])
	fmt.Println("last user:", users[len(users)-1])

	for i := 0; i < len(users); i++ {
		fmt.Println(i, users[i])
	}
}
