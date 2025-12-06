package main

import "fmt"

// To initialize a variable with a value, attribute the values in the same line.
var i, j int = 1, 2

func main() {
	// If the variables have a initilizer, the type can be inferred and omited
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

