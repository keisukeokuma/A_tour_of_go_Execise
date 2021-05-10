package main

// picパッケージを取得
import "golang.org/x/tour/pic"

// Pic関数を定義 引数はdx,dyで共にint型 返り値はuint8型の二次元スライスとする
func Pic(dx, dy int) [][]uint8 {
	// make(型,長さ,容量) スライス作成 長さ：スライス内に実際に入る変数の個数 容量：スライスに入る変数の最大個数
	// スライスret uint8型の二次元スライス、長さdx、容量dyとする
	ret := make([][]uint8, dx, dy)
	for y := 0; y < dy; y++ {
		// スライスrow uint8型のスライス、長さdxとする
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// row内の変数にuint8(x + y)の結果を代入する
			row[x] = uint8(x + y)
		}
		// retにrowを代入する
		ret[y] = row
	}
	return ret
}

func main() {
	// picパッケージ内のShow関数を実施し、引数を関数Picとする。
	// この時、関数Picの引数はpicパッケージ内に記載されており、dx=256,dy=256となる
	pic.Show(Pic)
}
