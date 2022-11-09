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

### Array in Go
```go
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
```

### Range in Go
we can iterate over an array using range also
```go
users := []string{
  "Khadga Shrestha", "Nishuka Shrestha", "Menuka Shrestha",
}

for i, name := range users {
  fmt.Println(i, name)
}
```

### Structure in Go
in go,we can use structure to define and use multiple data type [similar to class]
```go
type User struct {
	name     string
	age      int
	isActive bool
	skills   []string
}

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
```

### Function to the Structure
we can extend the structure using function by defining a function over it
```go
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

//after that we can use showInfo function as the member of User structure
user := User{
  name:     "Khadga Bahadur Shrestha",
  age:      25,
  role:     "Mobile Application Developer",
  isActive: true,
  skills: []string{
    "Flutter", "Node Js", "Docker", "MySQL",
  },
}

fmt.Println(user.showInfo())
```


### Map in Go
map is data structure to hold key-value pair
```go
userDoc := map[string]string{
		"name":     "Khadga Bahadur Shrestha",
		"role":     "Flutter Developer",
		"address":  "Madhumalla Morange",
		"username": "kodega2016",
}

fmt.Println(userDoc)
fmt.Println(userDoc["address"])
fmt.Println(userDoc["age"])

//we can check if value exists on the map as follows

age, isOk := userDoc["age"]
if isOk {
  fmt.Println("age:", age)
}
```


### Pointers in Go
```go
func main() {
	var p *int
	a := 10
	p = &a

	fmt.Println("Address:", p)
	fmt.Printf("Value with %d address is %d", p, *p)
}
```