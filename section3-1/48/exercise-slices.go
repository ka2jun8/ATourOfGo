package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8 = make([][]uint8, dy)
	for i := range result {
		x := make([]uint8, dx)
		result[i] = x
		for j := range result[i] {
			result[i][j] = uint8((j ^ i))
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}

// https://codebeautify.org/base64-to-image-converter
