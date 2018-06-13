package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("When's Saturday")
  // 曜日を取得する
  today := time.Now().Weekday()
  // switch文は上から下へcase評価します
  // caseの条件は一致すればそこで停止します（自動的にbreak）
  switch time.Saturday {
  case today + 0:
    fmt.Println("Today.")
  case today + 1:
    fmt.Println("Tomorrow.")
  case today + 2:
    fmt.Println("In two days")
  case today + 5:
    fmt.Println("立ち上がれ社畜共！月曜日だ！！")
  default:
    fmt.Println("Too far away.")
  }
}
