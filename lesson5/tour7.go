package main

import (
  "golang.org/x/tour/tree"
  "fmt"
)

// Walk walks the tre t sending all values
// from the tree to the channel ch
func Walk(t *tree.Tree, ch chan int) {
  var walker func(t *tree.Tree)
  walker = func(t *tree.Tree) {
    if t.Left != nil {
      walker(t.Left)
    }

    ch <- t.Value

    if t.Right != nil {
      walker(t.Right)
    }
  }

  walker(t)
  close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1,t2 *tree.Tree) bool {

  ch1 := make(chan int)
  ch2 := make(chan int)

  go Walk(t1,ch1)
  go Walk(t2,ch2)

  for value := range ch1 {
    if value != <-ch2 {
      return false
    }
  }
  return true
}

func main() {
  ch := make(chan int)
  go Walk(tree.New(1), ch)
  for n := range ch {
        fmt.Println(n, ",")
    }
  fmt.Println(Same(tree.New(1), tree.New(1)))
  fmt.Println(Same(tree.New(1), tree.New(2)))
}



/*
参考
http://www.geocities.jp/m_hiroi/golang/abcgo10.html
https://gist.github.com/takatoshiono/6835c961892497ac9fd28f7710cf22b1
*/
