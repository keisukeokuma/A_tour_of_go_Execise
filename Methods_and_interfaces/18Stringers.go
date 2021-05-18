package main

import "fmt"

/*
 * 学習内容
 * ・インターフェイスはメソッドを複数記述し、リスト化する。
 * ・メソッドはレシーバ引数を取る関数。typeで型定義されたものをレシーバ引数として取る。
 */

// byte型で4個の変数が格納できる配列IPAddrを作成
type IPAddr [4]byte

// インターフェイスStringerの中のメソッドStringを実装する
// IPAddr型の内容をIPアドレス(v4)の表現に直した文字列を返却する
func (a IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}

func main() {
	// mapの作成 キーはstring型、値はIPAddrでとする。
	hostMap := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	// 繰り返し処理を実施。キーをk、値をvとして出力する。fmt.Plintln実施時にString()も実施される。
	for _, v := range hostMap {
		fmt.Println(v)
	}
}
