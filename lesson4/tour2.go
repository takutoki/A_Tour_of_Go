package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

// ""メソッド""は、レシーバ引数を伴う""関数""ということです
// 一つ前のレッスンで実装した機能を変えずに通常の関数として記述できます
func Abs(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3,4}
  fmt.Println(Abs(v))
}
