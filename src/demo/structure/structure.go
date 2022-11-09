package main

import "fmt"

func main() {
	user := User{
		name:     "Khadga Bahadur Shrestha",
		age:      25,
		role:     "Mobile Application Developer",
		isActive: true,
		skills: []string{
			"Flutter", "Node Js", "Docker", "MySQL",
		},
	}

	fmt.Println(user)
	fmt.Println(user.name)
	fmt.Println(user.skills)

	fmt.Println(user.showInfo())

}

type User struct {
	name     string
	age      int
	role     string
	isActive bool
	skills   []string
}

func (user User) showInfo() string {
	return fmt.Sprintf("%s is %s aged %d\n", user.name, user.role, user.age)
}
