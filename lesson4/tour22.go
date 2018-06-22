package main

import (
  "fmt"
  "io"
  "strings"
  "golang.org/x/tour/reader"
)

type MyReader struct{}

func (r *MyReader) Read(b []byte) (n int, err error) {
  for {
    copy(b, "A")
    return 1,nil
  }
}

func main() {
  reader.Validate(MyReader{})
}
