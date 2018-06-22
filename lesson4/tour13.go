package main

import "fmt"

type I interface {
  M()
}

func main() {
  // インターフェース自体には値も具体的な型も保持していません
  // 具体的な型を代入しないままインターフェースを呼び出すとランタイムエラーとなります
  var i I
  describe(i)
  i.M()
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}

// 実行結果
// panic: runtime error: invalid memory address or nil pointer dereference
