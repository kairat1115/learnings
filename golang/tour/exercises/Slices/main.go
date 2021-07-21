package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for x := range arr {
		arr[x] = make([]uint8, dx)
	}
	for x := range arr {
		for y := range arr[x] {
			tmp := x + y
			switch {
			case tmp%12 == 0:
				arr[x][y] = uint8((x + y) / 2)
			case tmp%5 == 0:
				arr[x][y] = uint8(x * y)
			case tmp%3 == 0:
				arr[x][y] = uint8(x ^ y)
			}
		}
	}
	return arr
}

func main() {
	pic.Show(Pic)
}
