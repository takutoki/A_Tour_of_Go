package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

// メソッドは特別なレシーバ引数を関数に取ります
// funcキーワードと Abs()関数名の間に自身の引数リストで表現します
// この例では、Absメソッドは v という名前の Vertex型のレシーバを持つことを意味します
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3,4}
  fmt.Println(v.Abs())
}
