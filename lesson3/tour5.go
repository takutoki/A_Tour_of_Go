package main

import "fmt"

type Vertex struct {
  x, y int
}

var (
  v1 = Vertex{1, 2} // Vertexストラクト
  v2 = Vertex{x: 1} // y は暗黙的に 0 となる
  v3 = Vertex{}     // x,y ともに暗黙的に 0 となる
  p = &Vertex{1, 2} // Vertexストラクトへのポインタ
)

func main () {
  fmt.Println(v1, v2, v3, *p)
}
