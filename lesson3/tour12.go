package main

import "fmt"

func main() {
  // スライスのゼロ値は nil です
  var s []int
  fmt.Println(s, len(s), cap(s))
  if s == nil {
    fmt.Println("nil")
  }
}

/*
nil スライスは０の長さ(len)と容量(cap)をもち
何の元となる配列を持っていない
*/
