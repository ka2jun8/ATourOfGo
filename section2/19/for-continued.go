package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { // vscodeに勝手に最適化された
		sum += sum
	}
	fmt.Println(sum)
}
