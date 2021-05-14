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

// Errorメソッドを実装する
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("負の値のため計算できません:%v", float64(e))
}

// Sqrt関数を定義 引数はfloat64型、返り値はfloat64型とerrorとする。
func Sqrt(x float64) (float64, error) {
	// xが負の値であったときに、errorにnil以外の値が入るようにする。
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
	// Sqrt関数を実行
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
