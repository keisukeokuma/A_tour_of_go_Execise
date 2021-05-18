package main

import (
	"fmt"
	"math"
)

/*
 * errorインターフェイスについて
 * errorインターフェイスは返り値errorにnil以外が入っていた場合にError()メソッドの内容を実行します。
 */

// float64型を持つErrNegativeSqrt型を新たに作成。
type ErrNegativeSqrt float64

// Errorメソッドを実装する。
func (e ErrNegativeSqrt) Error() string {
	// eをfloat64型に変換してfmt.Sprintfで出力する。
	return fmt.Sprintf("負の値のため計算できません:%v", float64(e))
}

// Sqrt関数を定義 引数はfloat64型、返り値はfloat64型とerrorインターフェイスを満たすErrNegativeSqrt型を返す
func Sqrt(x float64) (float64, error) {
	// xが負の値であったときに、ErrNegativeSqrt型に型変換をしたxを返す。
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	// 28-43行目はflow_control_08の内容をコピーしたもののため、説明を省略します。
	z := 1.0
	delta := 1.0
	min_delta := 0.00000000001
	i := 0

	for ; math.Abs(delta) > min_delta; i++ {
		delta = (z*z - x) / (2 * z)
		z = z - delta
	}

	if i > 10 {
		fmt.Printf("反復回数は%d回で10回以上です\n", i)
	} else {
		fmt.Printf("反復回数は%d回で10回未満です\n", i)
	}

	// xが負の値でなかった場合、errorはnilを返す。
	return z, nil
}

func main() {
	// casesスライスを作成
	cases := []float64{2, -2}
	for i, v := range cases {
		// i+1で繰り返されている回数を表示
		fmt.Printf("%d 回目。＝＝＝＝＝\n", i+1)
		var ret float64
		var err error
		// err にSqrt(v)の結果を代入し、errがnilであった場合にif文内の処理を実施
		if ret, err = Sqrt(v); err != nil {
			// Error()の戻り値を表示
			fmt.Println(err)
			continue
		}
		// retを表示
		fmt.Println(ret)
	}
}
