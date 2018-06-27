package main

import "fmt"

func main() {
  // チャネルはバッファとして使えます
  // バッファを持つチャネルを初期化するのは、make関数の２つ目の引数にバッファの長さを与えます
  // 以下の例では、２つのバッファををもつため、int型の値を２つまで与えることができます
  // チャネルの値の出し入れはキュー
  ch := make(chan int, 2)
  ch <- 1
  ch <- 2
  fmt.Println(<-ch)
  fmt.Println(<-ch)
}
