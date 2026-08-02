package main

import "fmt"

func main() {
	// 初期化と後処理ステートメントの記述は任意
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
