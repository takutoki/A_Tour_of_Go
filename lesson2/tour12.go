package main

import "fmt"

func main() {
  // deferステートメントは、deferへ渡した関数の実行を、
  // 呼び出し元関数の終わり（returnする）まで遅延させるもの
  // defer を評価したタイミングの値が保持される
  world := "america"
  defer fmt.Println(world)

  world = "japan"
  fmt.Println("hello")
  fmt.Println(world)
  fmt.Println(".")
  fmt.Println("..")
  fmt.Println("...")
  fmt.Println("....")
  fmt.Println(".....")
  fmt.Println("......")
}
