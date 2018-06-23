package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	a := float64(x + y)
	b := uint(f)
	fmt.Println(x, y, z)
	fmt.Println(a, b)
}
