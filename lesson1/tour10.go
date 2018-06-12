package main

import "fmt"

func main() {
  var i, j int = 1, 2
  // 関数の中では、var 宣言の代わりに短い[:=]の代入文を使い暗黙的な型宣言ができる
  // 関数の外では[:=]を使うことはできない var func　などを使いましょう
  k := 3
  c, python, java := true, false, "no!"

  fmt.Println(i, j, k, c, python, java)
}
