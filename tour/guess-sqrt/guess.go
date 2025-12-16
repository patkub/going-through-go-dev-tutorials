package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	guesses := 15

	// A decent starting guess for z is 1, no matter what the input.
	z := 1.0

	for i := 0; i < guesses; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
