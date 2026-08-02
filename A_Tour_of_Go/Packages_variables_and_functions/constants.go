package main

import "fmt"

const Pi = 3.14

func main() {
	// 定数の宣言は`const`キーワードを使う
	// := は使えない
	// 文字、文字列、数値、bool値のみ定数化可能
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
