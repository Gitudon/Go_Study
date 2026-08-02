package main

import "fmt"

func main() {
	// 型を明示しない場合、型推論が行われる
	v := "change me" // change me!
	fmt.Printf("v is of type %T\n", v)
}
