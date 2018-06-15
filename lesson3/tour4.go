package main

import "fmt"

type Vertex struct {
  x int
  y int
}

func main() {
  v := Vertex{1, 2}
  p := &v
  // structのポインタを通してフィールドxにアクセスするには
  // (*p).x のように書くことができます。
  // しかしこの表記法は大変面倒ですので、Goでは代わりに p.x と書くこともできます。
  p.x = 1e9
  fmt.Println(v)
}
