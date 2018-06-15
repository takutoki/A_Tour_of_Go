package main

import "fmt"

func main() {
  // Arrays の宣言
  // 以下は stringの2つの配列を宣言
  // 宣言後は配列の長さを変えることはできません
  var a [2] string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

// 配列の宣言と初期化を一度に行う
  primes := [6]int{2,3,5,7,11,13}
  fmt.Println(primes)
}
