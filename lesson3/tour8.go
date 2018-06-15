package main

import "fmt"

func main() {
  // 固定長配列
  names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
  }
  fmt.Println(names)

  // スライスで取り出す
  a := names[0:2]
  b := names[1:3]
  fmt.Println(a, b)

  // スライスの要素を変更する
  b[0] = "xxx"
  // スライスは配列への参照のようなものです
  // スライスはどんなデータも格納しておらず、元の配列の部分列を指し示しています
  // スライスの要素を変更すると、その元となる配列の対応する要素も変更されます
  fmt.Println(a, b)
  fmt.Println(names)
}
