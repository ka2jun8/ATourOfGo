package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	v1, ok1 := m["Answer"]
	fmt.Println("The value: ", v1, "Present?", ok1)

	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"])

	v2, ok2 := m["Answer"]
	fmt.Println("The value: ", v2, "Present?", ok2)
}
