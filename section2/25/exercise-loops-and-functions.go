package main

import (
	"fmt"
	"math"
)

// Sqrt is a method calculated sqrt of x
func Sqrt(x float64) (float64, float64) {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, math.Sqrt(x)
}

func main() {
	fmt.Println(Sqrt(3))
}
