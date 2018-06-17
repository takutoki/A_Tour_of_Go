package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
  	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := &Vertex{3, 4}
  fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

  v.Scale(5)
  fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}


/*
ポインタレシーバを使う２つの理由
１．メソッドがレシーバが指す先の変数値を変更するため
２．メソッドの呼び出し毎に変数のコピーをを作るのを避けるためです
　　例えばレシーバが大きな構造体である場合有効です

注意点として
一般的に、変数レシーバ、または、ポインタレシーバのどちらかで
すべてのメソッドを与え、混在させるべきではありません

*/
