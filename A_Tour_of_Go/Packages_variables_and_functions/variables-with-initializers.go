package main

import "fmt"

// 初期化子を与えることができる
var i, j int = 1, 2

func main() {
	// 初期化子が与えらえたら型を省略可能で、初期化子が持つ型となる
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
