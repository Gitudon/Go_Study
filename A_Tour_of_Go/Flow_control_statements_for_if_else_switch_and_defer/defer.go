package main

import "fmt"

func test() {
	defer fmt.Println("test end")
	fmt.Println("test start")
}

func main() {
	// 呼び出し元の関数が終わるまで実行を遅延する
	defer fmt.Println("world")
	fmt.Println("hello")
	test()
}
