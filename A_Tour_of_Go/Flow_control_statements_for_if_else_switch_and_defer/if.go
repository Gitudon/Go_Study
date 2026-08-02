package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// 条件に()は不要、{}は省略不可
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
