package main

import "fmt"

func main() {
	var name string = "Khadga Bahadur Shrestha"
	fmt.Println(name)
	var age int = 24
	fmt.Println(age)

	username, role := "kodega2016", "Flutter Developer"
	fmt.Println(username, role)

	var (
		companyName = "Port Pro Nepal"
		address     = "Lalitpur Nepal"
	)

	fmt.Println(companyName, address)
}
