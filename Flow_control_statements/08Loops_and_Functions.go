package main

// fmtパッケージを取得
import (
	"fmt"
	"math"
)

// Sqrt関数を定義 引数はfloat64型、返り値はfloat64型とする。
func Sqrt(x float64) float64 {
	// zの初期値を宣言 float64型とするために1.0とする
	z := 1.0
	delta := 1.0
	// 最低値を定義
	min_delta := 0.00000000001
	// for文で繰り返した回数を出力するためにiをfor文外で定義
	i := 0

	// for文で変化がmin_deltaを下回るまで繰り返す。
	for ; math.Abs(delta) > min_delta; i++ {
		// 平方根の計算 問題文より引用
		delta = (z*z - x) / (2 * z)
		z = z - delta
	}

	// if文で回数と10回以上か未満かを出力する
	if i > 10 {
		fmt.Printf("反復回数は%d回で10回以上です\n", i)
	} else {
		fmt.Printf("反復回数は%d回で10回未満です\n", i)
	}
	return z
}

func main() {
	val := float64(2)
	// 関数の結果を出力
	fmt.Printf("自作関数の計算結果:%v \n", Sqrt(val))
	// math.Sqrtの結果を出力
	fmt.Printf("math.Sqrtの計算結果:%v \n", math.Sqrt(val))
}
