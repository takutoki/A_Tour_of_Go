package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := float64(1.0)
  s := float64(0)
  for i := 0; i < 10; i++ {
    z -= (z*z -x) / (2*z)
    // 1e-10は指数表現で、「1 × 10 の -10乗」という意味
    if math.Abs(z - s) < 1e-10 {
      break
    }
  }
  return z
}

func main() {
  fmt.Println(Sqrt(2))
}
