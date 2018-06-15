package main

import "fmt"

// struct は構造体で、フィールドの集まりです
type Vertex struct {
  x int
  y int
}

func main() {
  fmt.Println(Vertex{1, 2})
}
