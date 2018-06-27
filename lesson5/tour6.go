package main

import (
  "fmt"
  "time"
)

func main() {
  tick := time.Tick(100 * time.Millisecond)
  boom := time.After(500 * time.Millisecond)

  fmt.Printf("tick:%T boom:%T\n", tick, boom) // tick:<-chan time.Time boom:<-chan time.Time

  // どのcaseも準備できていないのであれば、selectの中のdefautが実行されます
  // ブロックせずに送受信するなら、defaultのcaseを使ってください
  for {
    select {
    case <-tick:
      fmt.Println("tick.")
    case <-boom:
      fmt.Println("BOOM!")
      return
    default:
      fmt.Println("    .")
      time.Sleep(50 * time.Millisecond)
    }
  }
}
