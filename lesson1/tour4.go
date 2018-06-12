package main

import "fmt"

// 関数の定義
// 変数名の後ろに型名を書く
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
