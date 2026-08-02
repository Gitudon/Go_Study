package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// 型変換を行える
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
