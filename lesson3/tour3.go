package main

import "fmt"

type Vertex struct {
  x int
  y int
}

func main() {
  v := Vertex{1, 2}
  // structフィールドにアクセスするにはドット（.）を用いる
  v.x = 4
  fmt.Println(v.x)
}
