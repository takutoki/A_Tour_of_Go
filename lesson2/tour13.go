package main

import "fmt"

func main() {
  fmt.Println("counting")

  // defer で渡した関数が複数ある場合、スタックされます
  // LIFOで実行されます
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("done")
}
/* run...
counting
done
9
8
7
6
5
4
3
2
1
0
*/
