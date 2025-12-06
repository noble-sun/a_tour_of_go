package main

import "fmt"

// This syntax is saying that the function 'add' have to take two 'int' parameters
// and it will also return an 'int'
func add(x int, y int) int {
	return x + y
}

func main() {
	// The order in which 'add' and 'main' is defined does not matter when calling
	// the functions.
	fmt.Println(add(42, 13))
}

