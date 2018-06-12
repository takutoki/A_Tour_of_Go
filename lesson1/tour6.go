package main

import "fmt"

// 関数は複数の戻り値を返せる
func swap(x, y string) (string, string) {
  return y, x
}

func main() {
  // 後のステップで出てくるが、
  // [:=]は変数宣言と初期化を一行で行うシンタックス
  a, b := swap("hello", "world")
  fmt.Println(a, b)
}
