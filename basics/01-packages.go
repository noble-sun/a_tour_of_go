package main

// By convention, take the last element of the import path. So "math/rand"
// is called using 'rand'
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println("This is another number:", rand.Intn(10))
}

// A GO program is made up of packages, and every program start running in the
// package main.
