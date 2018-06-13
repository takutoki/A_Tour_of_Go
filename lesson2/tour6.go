package main

import (
  "fmt"
  "math"
)

func pow(x, n, lim float64) float64 {
  // math.Pow関数は、ベースxをn値で累乗した値を返します
  // if文もfor文同様に、セミコロンで区切って初期処理がかける
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}

func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}
