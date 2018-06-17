package main

import (
  "fmt"
  "math"
)

type Abser interface {
  Abs() float64
}

func main() {
  var a Abser
  f := MyFloat(-math.Sqrt2)
  v := Vertex{3, 4}

  a = f   // インターフェイスのAbserの変数に、MyFloatを代入することで継承となる
  a = &v  // インターフェイスのAbserの変数に、*Vertexを代入することで継承となる

 // このラインでエラーとなる
 // Abs
  a = v

  fmt.Println(a.Abs())
}

type MyFloat float64

type Vertex struct {
  X, Y float64
}

// 自分で定義した MyFloat型 に対してインターフェイスの実装
func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

// 自分で定義した Vertex型 に対してインターフェイスの実装
func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
