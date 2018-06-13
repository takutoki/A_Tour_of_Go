package main

import "fmt"

func main() {
  sum := 1
  // セミコロンも省略できるのでwhile文のようにかける
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}
