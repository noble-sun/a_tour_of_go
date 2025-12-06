package main

import "fmt"

// The expected return type for functions that return multiple values 
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// This is using the short declaration := to assign the return values os the
	// swap/2 function to the variables 'a' and 'b' and infer the type of those vars.
	// Using the simple variable declaration would look something like this:
	//		var a, b int
	//		a, b = swap("hello", 100)
	// This would not work because 'go' ins't inferring the type anymore 
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

