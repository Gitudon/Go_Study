package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	// 初期化子を与えないと、ゼロ値が与えられる
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
