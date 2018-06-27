package main

import (
  "fmt"
  "time"
)

func say(s string) {
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

func main() {
  // goroutine はGoランタイムに管理される軽量なスレッドです
  go say("world")
  say("Hello")
}
