package main

import "fmt"

func sum(s []int, c chan int) {
  sum := 0
  for _, v := range s {
    sum += v
  }
  // c チャネルに変数 sum を送信する
  // もし以下の実行タイミングになってもmain内で
  // チャネルの準備できていない場合は送受信がブロックされます
  // これにより明確なロックや条件変数がなくてもgoroutineの同期が可能となります
  c <- sum
}

func main() {
  s := []int{7, 2, 8, -9, 4, 0}

  // チャネルの宣言　int型を受け取ります
  c := make(chan int)
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2:], c)
  // c チャネルから値を受け取る
  x, y := <-c, <-c

  fmt.Println(x, y, x+y)
}
