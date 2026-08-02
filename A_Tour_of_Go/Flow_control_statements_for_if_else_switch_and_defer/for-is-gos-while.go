package main

import "fmt"

func main() {
	sum := 1
	// セミコロンは省略可
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
