package main

import "fmt"

func main() {
	var i, j int = 1, 2

	// The short variable declaration symbol ':=' will declare and infer the type
	// of the variable.
	// It can only be used as a local variable. Outside of function variables have
	// to use the 'var' keyword.
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

