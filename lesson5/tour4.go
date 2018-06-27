package main

import "fmt"

func fibonacci(n int, c chan int) {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    c <- x
    x, y = y, x+y
  }
  // 送り手は、これ以上の送信する値がないことを示すために
  // チャネルを close できます
  // 受け手は、受信の式に２つ目のパラメーターを割当てることでそのチャネルがcloseされているかどうか確認できます
  // v, ok := <-ch
  close(c)
}

func main() {
  c := make(chan int, 10)
  go fibonacci(cap(c), c)
  // for i := range c はチャネルが閉じられるまで、チャネルのから値を繰り返し受信し続けます
  for i := range c {
    fmt.Println(i)
  }
}

/*
注意
送り手のチャネルだけをcloseしてください
受け手はcloseしてはいけません
もしcloseしたチャネルへ送信すると、panicします

チャネルはファイルオープンとは異なり、通常はcloseする必要はありません
closeするのはこれ以上値が来ないことを受け手が知る必要があるときだけです
*/
