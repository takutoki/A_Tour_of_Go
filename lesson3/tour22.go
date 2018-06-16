package main

import "fmt"

func main() {
  m := make(map[string]int)

  // map への操作
  // m へ要素の挿入や更新
  // m[key] = element
  // m の要素の取得
  // element = m[key]
  m["Answer"] = 42
  fmt.Println("The value:", m["Answer"])

  m["Answer"] = 48
  fmt.Println("The value:", m["Answer"])

  // m の要素の削除
  // delete(m, key)
  delete(m, "Answer")
  fmt.Println("The value:", m["Answer"])

  // キーに対する要素が存在するかどうかは、２つ目の返り値で確認します
  // element, ok = m[key]
  // m に key があれば変数 ok は true になります　なければ false です
  // m に key が存在しない場合 element はゼロ値になります
  v, ok := m["Answer"]
  fmt.Println("The value:", v, "Present?", ok)

  m["NO"] = 30

  nv, nok := m["NO"]
  fmt.Println("The value:", nv, "Present?", nok)
}
