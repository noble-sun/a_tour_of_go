package main

import (
	"fmt"
	"math"
)

// Exported names begin with a capital letter. 'Pi' is exported from the 'math'
// package, so we can access it using the 'importedPackage.ExportedPackage' convention
func main() {
	fmt.Println(math.Pi)
}

