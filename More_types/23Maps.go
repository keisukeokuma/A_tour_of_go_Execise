package main

// stringsパッケージ、wcパッケージを取得
import (
	"strings"

	"golang.org/x/tour/wc"
)

/**
* wordCount関数を定義
* 引数はstring型のsとし、返り値はキーはstring型、値はint型のmapを返す。
 */
func wordCount(s string) map[string]int {

	/**
	 * strings.Fields関数により、引数で渡された文章を空白文字ごとに分割した単語をスライスとして返す
	 * 空白のみが含まれる場合は空のスライスを返す
	 */
	word_list := strings.Fields(s) // ex) "  a b cc" -> []string{"a", "b", "cc"}

	// mapを作成する。キーはstring型、値はint型を指定する(デファルト値は0)。
	words := make(map[string]int, len(word_list))

	// rangeで繰り返し処理を実施。word_listの値のコピーを変数wに定義する。
	for _, w := range word_list {
		// スライスword_listから単語を取り出し、wordsのキーとする。同時に値をインクリメントする。
		// mapは重複を許さないため、同じ単語が出た場合は既存の同じ単語のキーの値にインクリメントする
		words[w]++
	}
	return words
}

func main() {
	/**
	 * wc.Test関数内にはwordCountの計算対象となる文字列と単語の出現回数を計算した正答が記載されている。
	 * 正しく単語の出現回数を計算出来た場合は"PASS"に続いて、文字列と計算結果を出力する。
	 * 計算が失敗した場合は、"FAIL"に続いて計算結果と正答が出力される
	 */
	wc.Test(wordCount)
}
