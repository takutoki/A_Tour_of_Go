package main

import "fmt"

func main() {
  sum := 1
  // for文の初期化子と後処理ステートメントの記述は任意のためなくてもよい
  fot ; sum < 1000 ; {
      sum += sum
  }
}
