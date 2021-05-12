package main

// picパッケージを取得
import "golang.org/x/tour/pic"

// Pic関数を定義 引数はdx,dyで共にint型 返り値はuint8型の二次元スライスとする
func Pic(dx, dy int) [][]uint8 {
	// make関数でスライスを作成する。make(型,長さ,容量)
	// 二次元スライスret uint8型の二次元スライス、長さ(スライス内の要素数)はdx、容量(スライスに入る要素の最大個数)はdyとする
	ret := make([][]uint8, dx, dy)
	for y := 0; y < dy; y++ {
		// スライスrow uint8型のスライス、長さdxとする
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// スライスrowのx番目の要素にuint8(x + y)の結果を代入する
			// uint8(x + y)の結果は、0-255までの数値が入る
			row[x] = uint8(x + y)
		}
		// 二次元スライスretのy番目の要素にrowを代入する
		ret[y] = row
	}
	return ret
}

func main() {
	/**
	* picパッケージのShow関数を呼び出す。引数は関数をとり、今回はPic関数を引数とする。
	* Pic関数で定義されているint型のdx,dyは、Show関数内で定数256がそれぞれ代入される。
	 */
	pic.Show(Pic)
}
