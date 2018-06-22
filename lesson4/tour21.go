package main

import (
  "fmt"
  "io"
  "strings"
)

func main() {
  r := strings.NewReader("Hello, Reader!")
  fmt.Printf("%T\n", r)   // *strings.Reader

  b := make([]byte, 8)
  for {
    // データを与えられたバイトスライスへ入れ
    // 入れたバイトのサイズとエラーの値を返します
    // ストリームの終端の場合は err に io.EOF を返します
    n, err := r.Read(b)
    fmt.Printf("n = %v err = %v b = %v", n, err, b)
    fmt.Printf("b[:n] = %q\n", b[:n])
    if err == io.EOF {
      break
    }
  }
}
