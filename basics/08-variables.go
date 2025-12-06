package main

import "fmt"

// 'var' declares a single or multiple variables, and its type.
// This can be declared as a global package variable level, or as a function level.
var c, python, java bool

func main() {
	// 'i' will only be accessable inside the 'main' function, but the variables
	// declared outside the function can be accessed inside any function.
	var i int
	fmt.Println(i, c, python, java)
}

