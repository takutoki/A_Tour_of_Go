package main

import (
  "fmt"
  "math"
)

func pow(x, n, lim float64) float64 {
  // if文で宣言された変数はelseブロック内でも使うことができます
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %G\n", v, lim)
  }
  // can't use v here, though
  return lim
}

func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}
