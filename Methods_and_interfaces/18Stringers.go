package main

import "fmt"

// byte型で4個の変数が格納できる配列IPAddrを作成
type IPAddr [4]byte

// インターフェイスを実装する
// 配列IPAddrにインターフェイスStringを実装する
func (a IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}

func main() {
	// mapの作成 キーはstring型、値はIPAddrを通したstring型
	hosts := map[string]IPAddr{
		"loopback":  {12, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// Interfaseはメソッドの定義を集める
//
