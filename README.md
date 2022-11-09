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

