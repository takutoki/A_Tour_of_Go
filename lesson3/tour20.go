package main

import "fmt"

// 構造体
type Vertex struct {
  Lat, Long float64
}

// map リテラルは、ストラクトに似ていますが、キーが必要です
var m = map[string]Vertex{
  "Bell Labs": Vertex{
    40.68433, -74.39967,
  },
  "Google": Vertex{
    37.42202, -122.08408,
  },
}

func main() {
  fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}
