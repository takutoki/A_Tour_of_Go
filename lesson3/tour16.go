package main

import "fmt"

var pow = []int{1, 2, 4 , 8 , 16, 32 , 64, 128}

func main() {
  // for ループに利用する range はスライスやマップ（map）をひとつずつ反復処理するために使います
  // スライスをrangeで繰り返す場合、反復毎に２つの変数を返します
  // 1つめは変数のインデックス、2つめはインデックスの場所の要素のコピーです 
  for i,v := range pow {
    fmt.Printf("2**%d = %d\n", i ,v )
  }
}
