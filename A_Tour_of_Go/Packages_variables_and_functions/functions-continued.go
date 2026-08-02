package main

import "fmt"

// 2つ以上の型名が同じなら、最後の型を残し省略して記述できる
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
