package main

import (
  "fmt"
  "math"
)

// 他の変数のように関数を渡すことができます
func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}

func main() {
  // 関数も変数です
  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }
  fmt.Println(hypot(5,13))

  fmt.Println(compute(hypot))
  fmt.Println(compute(math.Pow))
}
