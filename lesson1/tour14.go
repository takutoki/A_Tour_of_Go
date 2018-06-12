package main

import "fmt"

func main() {
  // 明示的な型を指定せずに変数宣言する場合、右辺から型推論されます
  // v の値を変えると型も変わることが確認できます
  v := 42
  fmt.Printf("v is of type %T\n", v)
}
