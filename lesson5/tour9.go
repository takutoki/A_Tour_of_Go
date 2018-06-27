package main

import (
  "fmt"
  "sync"
  "time"
)

// SafeCounter is safe to use concurrently
// コンフリクトを避けるために、一度に１つのgoroutineだけが変数にアクセスできるようにする
// 排他制御と呼ばれ、このデータ構造を指す名前は、mutex です
type SafeCounter struct {
  v   map[string]int
  mux sync.Mutex
}

// Go の標準ライブラリは、排他制御をsync.Mutexと次の２つのメソッドで提供します
// Lockメソッド Unlockメソッド
func (c *SafeCounter) Inc(key string) {
  c.mux.Lock()
  c.v[key]++
  c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
  c.mux.Lock()
  defer c.mux.Unlock()
  return c.v[key]
}

func main() {
  c := SafeCounter{v: make(map[string]int)}
  for i := 0; i < 1000; i++ {
    go c.Inc("somekey")
  }

  time.Sleep(time.Second)
  fmt.Println(c.Value("somekey"))
}
