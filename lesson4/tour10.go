package main

import "fmt"

// インターフェース
type I interface {
  M()
}

// 自身で定義した型Tに対してメソッドを実装していくことによってインターフェースを満たします
type T struct {
  S string
}

// このメソッドは、インターフェースIを継承している型Tを表しています
// しかしそれを明示的に宣言する必要はありません
func (t T) M() {
  fmt.Println(t.S)
}

func main() {
  var i I = T{"Hello"}
  i.M()
}
