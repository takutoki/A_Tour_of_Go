package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

// 変数レシーバメソッド
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
  	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(v.Abs())
  fmt.Println(AbsFunc(v))

  // メソッドが変数レシーバである場合、呼び出し時に
  // 変数、または、ポインタのいずれかをレシーバとして取ることができます
  p := &Vertex{4, 3}
  // p.Abs()ステートメントを
  // (*p).Abs() として解釈します
  fmt.Println(p.Abs())
  fmt.Println(AbsFunc(*p))
}
