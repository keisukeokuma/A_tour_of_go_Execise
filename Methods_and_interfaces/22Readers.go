package main

import (
	"golang.org/x/tour/reader"
)

// 構造体MyReaderを定義
type MyReader struct{}

// Readメソッドを定義する。値レシーバとしてrを受け取る。
// また、引数を[]byte型のbを取り、返り値をint型とerrorインターフェイスを満たす値を返す。
func (r MyReader) Read(b []byte) (int, error) {
	// 繰り返し処理
	for i := range b {
		b[i] = 'A'
	}
	// len(b)とnilを返す。 len(b)はbの長さを取得
	return len(b), nil
}

func main() {
	/*
	 * ValidateにReadメソッドを持つ構造体MyReader{}を引数として渡す。
	 * Validate内で'A'が繰り返し作成されているかをチェックし、成功であれば"OK"を出力する
	 */
	reader.Validate(MyReader{})
}
