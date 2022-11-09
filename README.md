# Go Fundamentals
Learn basic of go language 

Program Structure
like every other language in go there must be a main method which is the entry point of the application.
```go
package main

import "fmt"

func main() {
  fmt.Println("hello go")
}
```

## Variables in Go
We can define variable in go using var keyword as 
```go
var name string = "Khadga Bahadur Shrestha"
fmt.Println(name)
var age int = 24
```

In go we can define varibale using compound assignment also
```go
username, role := "kodega2016", "Flutter Developer"
fmt.Println(username, role)
```

we can also define variable using block assignment
```go
var (
		companyName = "Port Pro Nepal"
		address     = "Lalitpur Nepal"
)

fmt.Println(companyName, address)
```

## Operators in Go
like other languages, go has
- arithmetic operators
- assigmnet operators
- logical operators
- comparision operators
- bitwise operators

```go
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
```

## Functions in Go
in go,we can define function using `func` keyword


#### Basic function

```go
func greet() {
	fmt.Println("greeting...")
}
```

#### Function with arguments and return type
```go
func greetWithName(name string) string {
	return fmt.Sprintf("greeting...%s\n", name)
}
```

#### Function with multiple return
```go
func swap(x, y string) (string, string) {
	return y, x
}
```

## If-Else
```go
a, b := 20, 40

if a < b {
  fmt.Println("a is smaller than b")
} else {
  fmt.Println("a is greater than b")
}
```

### Switch case
```go
switch age := 16; {
case age == 0:
  fmt.Println("newborn")

case age > 0 && age <= 3:
  fmt.Println("toddler")

case age > 3 && age <= 12:
  fmt.Println("toddler")

case age > 3 && age <= 12:
  fmt.Println("child")

case age > 12 && age <= 17:
  fmt.Println("teenager")

case age > 18:
  fmt.Println("adult")
}
```