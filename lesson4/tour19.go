package main

import (
  "fmt"
  "time"
)

type MyError struct {
  When time.Time
  What string
}

/* error型はfmt.Stringerに似た組み込みのインターフェースです
type error interface {
  Error() string
}
fmt.Stringer と同様に、fmtパッケージは、変数を文字列で
出力する際に errorインターフェースを確認します
*/
// MyError構造体にError() 関数を実装して
// errorインターフェースを継承する
func (e *MyError) Error() string {
  return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
  return &MyError{time.Now(), "it didn't work"}
}

func main() {
  if err := run(); err != nil{
    fmt.Println(err)
  }
}
