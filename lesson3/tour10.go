package main

import "fmt"

func main() {
  s := []int{2,3,5,7,11,13}

  s = s[1:4]
  fmt.Println(s)

  // スライスするとき、low値を省略すると既定値を使用することになる
  // low の既定値は 0
  s = s[:2]
  fmt.Println(s)

 // high の既定値はスライスの長さ
  s = s[1:]
  fmt.Println(s)
}

/*
以下の配列において
var a[10]int

これらのスライスは等価です

a[0:10]
a[:10]
a[0:]
a[:]

*/
