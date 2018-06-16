package main

import "fmt"

type Vertex struct {
  // 経度　緯度
  Lat, Long float64
}

// map とはキーと値を関連付けた連想配列
// map のゼロ値は nil です
// nil マップはキーを持っておらず、またキーを追加することもできません
var m map[string]Vertex

func main() {
  // make 関数は指定された型を、初期化され使用できるようにして返します
  fmt.Println(m)                // map[] と出力
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
  }
  fmt.Println(m["Bell Labs"])
}
