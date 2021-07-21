package main

import (
	"fmt"
	"math"
)

func Abs(x float64) float64 {
	if x < 0.0 {
		return -x
	}
	return x
}

func Sqrt(x float64) float64 {
	z := 1.0
	prevZ := 0.0
	for {
		z -= (z*z - x) / (2 * z)
		if v := fmt.Sprintf("%f", Abs(z-prevZ)); v == "0.000000" {
			break
		}
		fmt.Println("CALCULATING:", z)
		prevZ = z
	}
	return z
}

func main() {
	n := float64(2)
	fmt.Println(Sqrt(n))
	fmt.Println(math.Sqrt(n))
}
