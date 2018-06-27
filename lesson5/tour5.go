package main

import "fmt"

// select ステートメントは、goroutineを複数の通信操作で待たせます
func fibonacci(c, quit chan int) {
  x, y := 0, 1
  for {
    // selsect ステートメントは複数あるcaseにのいずれかが準備できるようになるまでブロックし
    // 準備できたcaseを事項します
    // もし複数のcaseの準備ができている場合、caseはランダムに選択されます
    select {
    case c <- x:
      x, y = y, x+y
    case <-quit:
      fmt.Println("quit")
      return
    }
  }
}

func main() {
  c := make(chan int)
  quit := make(chan int)
  go func() {
    for i := 0; i < 10; i++ {
      fmt.Println(<-c)
    }
    quit <- 0
  }()
  fibonacci(c, quit)
}
