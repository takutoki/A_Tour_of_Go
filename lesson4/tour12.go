package main

import "fmt"

type I interface {
  M()
}

type T struct {
  S string
}

// インターフェース自体の中にある具体的な値がnilの場合
// メソッドは nil をレシーバとして呼び出します
func (t *T) M() {
  if t == nil {
    fmt.Println("<nil>")
    return
  }
  fmt.Println(t.S)
}

func main() {
  var i I

  var t *T
  i = t
  describe(i)
  // i に値が入っていないためnil
  i.M()

  i = &T{"hello"}
  describe(i)
  i.M()
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}
