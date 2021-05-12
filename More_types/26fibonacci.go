package main

import "fmt"

/**
* 本問題の全体説明(関数値・クロージャ)
* ・関数は変数のように関数に渡せる。引数にも戻り値にも出来る。
* ・関数fibonacci内で定義された無名関数は、fibonacci内のローカル変数にアクセス出来る。
* ・関数fibonacci内のローカル変数は、無名関数ごとに保存される
 */

// fibonacci関数の定義 戻り値をint型の値を返す無名関数とする
func fibonacci() func() int {
	// 変数定義 a,b共にint型で定義する
	a, b := 1, 0
	// 戻り値となる無名関数 int型の値を返す
	return func() int {
		// fibonacci内で定義したローカル変数a,bを使用し、計算を行ってa,bに再度代入される
		a, b = b, a+b
		return a
	}
}

func main() {
	// f変数を宣言し、fibonacci関数を代入する
	f := fibonacci() // => func f() float32{}
	for i := 0; i < 10; i++ {
		// fibonacci()を呼び出し、表示する。
		fmt.Println(f())
	}
}
