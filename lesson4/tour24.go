package main

import (
  "fmt"
  "image"
)

func main() {
  m := image.NewRGBA(image.Rect(0, 0 , 100, 100))
  fmt.Println(m.Bounds())
  fmt.Println(m.At(0, 0).RGBA())
}

/*
imageパッケージは以下のimageインターフェースを定義しています

type Image interface{
ColorModel() color.ColorModel
Bounds() Rectangle
At(x int) color.Color
}

返り値である color.ColorModel と colori.Color は共にインターフェースですが
定義済みの color.RGBA と color.RGBAModel を使うことで
このインターフェースは無視できます
これらのインターフェースは image/color パッケージで定義されています

*/
