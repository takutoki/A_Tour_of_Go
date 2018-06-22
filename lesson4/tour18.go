package main

import (
	"fmt"
	"strings"
	"strconv"
)

// byte型 は符号なし整数型8ビットです（uint8 の別名です）
type IPAddr [4]byte

func (ipa IPAddr) String() string {
	s := make([]string, len(ipa))
	for i, val := range ipa {
		//byte型からstring型へキャスト
		s[i] = strconv.Itoa(int(val))
	}
	return strings.Join(s, ".")
}

func main() {
  hosts := map[string]IPAddr{
    "lookback":  {127, 0, 0, 1},
    "googlesDNS": {8, 8, 8, 8},
  }
  for name, ip := range hosts {
    fmt.Printf("%v: %v\n", name, ip)
  }
}
