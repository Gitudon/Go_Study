package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// := varの代わりに使える暗黙的な型宣言
	// 関数の外ではvarを使わないといけない
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
