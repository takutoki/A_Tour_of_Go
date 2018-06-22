package main

import "fmt"

type Person struct {
  Name string
  Age int
}

// 構造体PersonはStringerクラスを継承している
func (p Person) String() string {
  // Sprintf関数は、指定されたフォーマットに従ってフォーマットを行い、結果を文字列で返します
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
  a := Person{"Arthur Dent", 42}
  b := Person{"Zaphod Beeblebrox", 9001}
  fmt.Println(a,b)
  // 構造体を初期化した[a,b]を出力する処理において
  // なぜfmt.Println(s.String(), b.String())ではなく、
  // 単純に (a, b) と呼び出しただけでなぜString()が呼ばれるのか分からない
  // 実際に a.String() というふうに記述した場合も結果は同じだった

  //解決→http://www.geocities.jp/m_hiroi/golang/abcgo09.html
  //fmt.Print/Printlnは、引数が Stringer 型であればメソッド String を呼び出します
  //Person型はStringer型を継承しているので、String()が呼ばれます
}

/*
わかりやすい説明！→　http://maku77.github.io/hugo/go/interface.html
Go 言語には明示的にインタフェース名を指定して実装することはありません。
ただ単純に、インタフェースによって示されているメソッドを実装するだけで、
その型はそのインタフェースを備えている（実装している）とみなされます。
通称ダックタイピングと言われているものです
（アヒルのように歩いて鳴けば、それはアヒルであるという考え方）。
上の例では、Person構造体は、Stirng()メソッドを実装しているので
GOコンパイラは fmt.Stringerインターフェースとしてみなします
*/
