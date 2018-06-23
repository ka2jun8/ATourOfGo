package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// Word count function
func WordCount(s string) map[string]int {
	result := map[string]int{}
	chars := strings.Fields(s)

	for _, char := range chars {
		_, ok := result[char]
		if ok == false {
			result[char] = 1
		} else {
			result[char]++
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
