package main

import "fmt"

var i, j int = 1, 2

func main() {
  // var 宣言では、変数ごとに初期化子を与えることができる
  // 初期化子が与えられている場合、型を省略できます
  // その変数は初期化子が持つ型になります
  var c, python, java = true, false, "no!"
  fmt.Println(i, j, c, python, java)
}
