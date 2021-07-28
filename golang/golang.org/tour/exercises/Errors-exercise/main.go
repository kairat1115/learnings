package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Abs(x float64) float64 {
	if x < 0.0 {
		return -x
	}
	return x
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return x, ErrNegativeSqrt(x)
	}

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
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
