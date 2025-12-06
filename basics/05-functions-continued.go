package main

import "fmt"

// If two or more consecutive named parameters have the same type, its possible
// to omit the type from all the others but the last one.
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

