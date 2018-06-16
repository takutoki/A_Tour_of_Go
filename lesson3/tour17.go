package main

import "fmt"

func main() {
  pow := make([]int, 10, 10)
  for i := range pow {
    pow[i] = 1 << uint(i) // 2のi乗を入れる
  }
  // インデックス値や実値は、"_"(アンダーバー)へ代入することによって捨てることができます
  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }
}
