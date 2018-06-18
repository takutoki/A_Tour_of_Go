package main

import "fmt"

func main() {
  var i interface{} = "hello"

  // 型アサーション
  // このシンタックスは、インターフェースの値iが具体的な型Tを保持し、
  // 基になるTの値を変数
  s := i.(string)
  fmt.Println(s)

  // インターフェースが特定の型を保持しているかどうかをテストするために
  // 型アサーションは２つの値を返すことができます
  // 基となる値と、アサーションが成功したかどうかを報告するブール値
  s, ok := i.(string)
  fmt.Println(s, ok)

  f, ok := i.(float64)
  fmt.Println(f, ok)

  // インターフェースが保持していない型を指定した場合panicを起こします
  f = i.(float64)    // panic
  fmt.Println(f, ok)
}
