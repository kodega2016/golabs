package main

import "fmt"

func main() {
	user := User{
		name:     "Khadga Bahadur Shrestha",
		age:      25,
		isActive: true,
		skills: []string{
			"Flutter", "Node Js", "Docker", "MySQL",
		},
	}

	fmt.Println(user)
	fmt.Println(user.name)
	fmt.Println(user.skills)

}

type User struct {
	name     string
	age      int
	isActive bool
	skills   []string
}
