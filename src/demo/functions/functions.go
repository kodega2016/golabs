package main

import "fmt"

func main() {
	greet()
	result := greetWithName("Nishuka shrestha")
	fmt.Println(result)
	sum := add(20, 20)
	fmt.Println("sum:", sum)
	fmt.Println(swap("Nepal", "Hello"))
}

func greet() {
	fmt.Println("greeting...")
}

func greetWithName(name string) string {
	return fmt.Sprintf("greeting...%s\n", name)
}
func add(a, b int) int {
	return a + b
}
func swap(x, y string) (string, string) {
	return y, x
}
