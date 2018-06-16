package main

import (
  "fmt"
)

func main() {
  var s []int
  printSlice(s)

  // append で nil スライスを追加
  s = append(s, 0)
  printSlice(s)

  // スライスは必要に応じて拡張する
  s = append(s, 1)
  printSlice(s)

  // 元配列sが、変数群を追加する際に容量（cap）が小さい場合
  // より大きいサイズの配列を割当て直す → 容量が増えたら別配列を指すということ
  // その場合、戻り値となるスライスは、新しい割当先を示すようになります
  s = append(s, 2, 3, 4)
  printSlice(s)
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
