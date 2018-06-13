package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Println("Go runs on")
  // switch文はif-else文を短く書くシンタックス
  // GOでは選択されたcaseだけを実行してそれに続くすべてのcaseは実行されません
  // breakステートメントが自動的に提供されます
  // またswitch の case は定数である必要はなく、関係する値は整数である必要はありません
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux")
  default:
    //その他OS
    fmt.Printf("%s.", os)
  }
}
