package main

import (
  "fmt"
  "math"
)

func sqrt(x float64) string {
  // if文は () は不要で {} は必要です（おなじみの書き方ですね！）
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func main() {
  fmt.Println(sqrt(2), sqrt(-4))
}
