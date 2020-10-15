package main

import "fmt"

func Increment(x int) (result int) {
	result = x + 1
	return result
}

func main() {
	fmt.Printf("Hi, %d is more than %d.", 2, Increment(2))
}
