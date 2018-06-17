package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

// レシーバ引数がある通常のメソッド
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインターレシーバでメソッドを宣言できます
// ポインターレシーバを持つメソッドは、レシーバが指す変数の値を変更できます
// レシーバ自身を更新することが多いため変数レシーバよりもポインタレシーバのほうが一般的です
func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3, 4}
  v.Scale(10)
  fmt.Println(v.Abs())
}

/*
レシーバメソッドは、構造体型（または構造体のポインタ型）にメソッドを設けることが可能
上記では、Vertex型を初期化した変数 v にScaleメソッドを実装したことになる
*/
