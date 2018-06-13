package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now()
  // 条件のないswitch文は switch true と書くこととと同じです
  // この場合のswitch文は長くなりがちな"if-then-else"のつながりをシンプルに表現できます
  switch  {
  case t.Hour() < 12:
    fmt.Println("Good Moning")
  case t.Hour() < 17:
    fmt.Println("Good afternoon")
  default:
    fmt.Println("Good evening")
  }
}
