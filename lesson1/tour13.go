package main

import (
  "fmt"
  "math"
)

func main() {
  // 型変換　明示的な変換が必要
  var x, y int = 3, 4
  // Sqrt関数は引数にfloat64型を受け取り平方根を返す
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Println(x, y, z)
}
