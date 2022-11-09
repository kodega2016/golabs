package main

import "fmt"

func main() {
	userDoc := map[string]string{
		"name":     "Khadga Bahadur Shrestha",
		"role":     "Flutter Developer",
		"address":  "Madhumalla Morange",
		"username": "kodega2016",
	}

	fmt.Println(userDoc)
	fmt.Println(userDoc["address"])
	fmt.Println(userDoc["age"])

	age, isOk := userDoc["age"]
	if isOk {
		fmt.Println("age:", age)
	}
}
