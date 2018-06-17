package main

import "fmt"

type Vertex struct {
  X, Y float64
}

// ポインターレシーバを受け取るメソッド
func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

// 引数にポインタを受け取る関数
func ScaleFunc(v *Vertex, f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3, 4}
  // メソッドがポインタレシーバの場合、呼び出し時に
  // 変数、または、ポインタのいずれかをレシーバとして取ることができます
  // 以下の v は変数でありポインタではありませんが、メソッドでポインタレシーバが自動的に呼び出されます
  // Goでは利便性のため、v.Scale(2) のステートメントを
  // (&v).Scale(2) として解釈します
  v.Scale(2)
  // ScaleFunc関数は引数としてポインタを受け取るので、ポインタを渡す必要があります
  ScaleFunc(&v, 10)

  p := &Vertex{4, 3}
  p.Scale(3)
  ScaleFunc(p, 8)

  fmt.Println(v, p)
}
