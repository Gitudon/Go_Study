package main

import "fmt"

// 戻り値に名前を付けることができるが、可読性は落ちる
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
