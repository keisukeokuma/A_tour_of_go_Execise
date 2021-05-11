package main

// picパッケージ、stringsパッケージを取得
import (
	"strings"

	"golang.org/x/tour/wc"
)

/**
* WordCount関数を定義
* 引数はstring型のsとし、返り値はキーはstring型、値はint型のmapを返す。
 */
func WordCount(s string) map[string]int {

	/**
	 * strings.Fields関数により、引数で渡された文字列を空白文字ごとに分割し、スライスとして返す
	 * 空白のみが含まれる場合は空のスライスを返す
	 */
	word_list := strings.Fields(s)

	// mapを作成する。キーはstring型、要素はint型を指定。
	words := make(map[string]int)

	// rangeで繰り返し処理を実施。word_listの要素のコピーを変数wに定義する。
	for _, w := range word_list {
		// wordsに分割した文字列(w)をキーとした要素を作成し、要素をインクリメントする
		// mapsは重複を許さないため、同じ単語が出た場合に既存の同じ文字列であるキーの要素にインクリメントする
		words[w]++
	}
	return words
}

func main() {
	/**
	 * wc.Test関数内にはWordCountの計算対象となる文字列と単語の出現回数を計算した正答が記載されている。
	 * 正しく単語の出現回数を計算出来た場合は"PASS"に続いて、文字列と計算結果を出力する。
	 * 計算が失敗した場合は、"FAIL"に続いて計算結果と正答が出力される
	 */
	wc.Test(WordCount)
}
