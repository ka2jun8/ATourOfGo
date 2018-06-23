package main

import "fmt"

func fibonacci() func() int {
	result, p1, p2, n := 0, 0, 1, -1

	return func() int {
		n++
		if n == 0 {
			return p1
		} else if n == 1 {
			return p2
		}
		result = p1 + p2
		p1 = p2
		p2 = result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
