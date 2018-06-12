package main

import "fmt"

// 戻り値の変数に名前をつける( named return value )ことができます
// 関数定義時に定義した変数として扱われる
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  // 名前をつけた戻り値の変数を使うと、returnステートメントに
  // 何も書かずに戻すことができます（"naked" return と呼ぶ）
  return
}

func main () {
  fmt.Println(split(17))
}
