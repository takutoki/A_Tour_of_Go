package main

import "fmt"

func main() {
  // スライスは組み込みの make関数で作成できます
  // これは動的サイズの配列を作成する方法です
  // 引数は、スライスの型、長さ、容量
  // 容量を省略すると、長さと同じになります
  a := make([]int, 5)
  printSlice("a", a)

  b := make([]int, 0, 5)
  printSlice("b", b)

  c := b[:2]
  printSlice("c", c)

  d := c[2:5]
  printSlice("d", d)

}

func printSlice(s string, x []int) {
  fmt.Printf("%s len=%d cap=%d %v\n", s, len(s), cap(x), x)
}
