package main

// fmtパッケージを取得
import (
	"fmt"
)

// Sqrt関数を定義 引数はfloat64型、返り値はfloat64型とする。
func Sqrt(x float64) float64 {
	// zの初期値を宣言 float64型とするために1.0とする
	z := 1.0
	for i := 0; i < 10; i++ {
		// 平方根の計算 問題文より引用
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	// 関数の結果を出力
	fmt.Println(Sqrt(2))
}
