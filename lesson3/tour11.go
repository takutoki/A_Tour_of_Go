package main

import "fmt"

func main() {
  s := []int{2,3,5,7,1,13}
  printSlice(s)

  // スライスに配列数０のスライスを与える
  s = s[:0]
  printSlice(s)

  // 配列数を拡張する
  s = s[:4]
  printSlice(s)

  // 2つだけ取り出す
  s = s[2:]
  printSlice(s)
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %V\n", len(s), cap(s), s)
}
