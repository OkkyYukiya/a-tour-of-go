package main

import (
	"fmt"
	"math"
)

func pow (x, n, lim float64) float64 {
	// vにmath.Pow(x, n)の結果が代入される
	// 続いてそのままvとlimを比較
	// trueであればvが返される
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		// ニュートン法を用いて平方根を求める 以下引用
		// コンピュータは通常ループを使って x の平方根を計算します。
		// いくつかの z を推測することから始めて、
		// z² がどれほど x に近づいているかに応じて z を調整できます。

		result := z - (z*z-x)/(2*z)

		// 収束条件として前回の値と一致したら解が出たと見て処理を終了させる楽しいなこれ
		if result == z {
			return z
		}
		z = result
		fmt.Println("i: ", i, "z: ", z)
	}
	return z
}

func main () {
	// fmt.Println(
	// 	pow(3,2,10),
	// 	pow(3,3,20),
	// )
	fmt.Println(Sqrt(2))
}