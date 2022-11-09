package main

import "fmt"

func main() {
	counter := Counter{
		count: 1,
	}

	fmt.Println("count:", counter.count)
	increment(&counter)
	fmt.Println("count:", counter.count)

}

type Counter struct {
	count int
}

func increment(counter *Counter) {
	counter.count++
}
