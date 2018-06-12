package main

import (
  "fmt"
  "math"
)

func main() {
  // 最初の文字が大文字で始まる名前は、外部のパッケージから参照できる
  // エクスポート（公開）された名前(下のPi)
  // mathパッケージをインポートすることで
  // パッケージがエクスポートしている名前を参照することができます
  fmt.Println(math.Pi)
}
