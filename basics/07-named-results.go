package main

import "fmt"

// The return values of a function can be named. The 'naked' return will return
// the named return values, so if the variables have something attributed to them
// inside the function, that value will be returned.
// If no value is attributed to the named return value, it will return the default
// value for the type, i believe.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

