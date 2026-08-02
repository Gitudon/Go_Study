package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	i := 0
	for {
		buf := z
		z -= (z*z - x) / (2 * z)
		i++
		if math.Abs(z-buf) < 0.000001 {
			break
		}
	}
	println(i)
	return z
}

func main() {
	fmt.Println(Sqrt(141))
	fmt.Println(math.Sqrt(141))
}
