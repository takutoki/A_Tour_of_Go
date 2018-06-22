package main

import "fmt"

func do(i interface{}) {
  // 型switch はいくつかの型アサーションを直列に使用できる構造です
  // 型switchは通常のswitch文と似ていますが、それらの型は指定された
  // インターフェースの値が保持する値の型を比較されます
  switch v := i.(type) {
  case int:
    fmt.Printf("Twice %v is %v\n", v, v*2)
  case string:
    fmt.Printf("%q is %v bytes long\n",v, len(v))
  default:
    fmt.Printf("I don't know about type %T!\n",v)
  }
}

func main() {
  do(21)
  do("hello")
  do(true)
}
