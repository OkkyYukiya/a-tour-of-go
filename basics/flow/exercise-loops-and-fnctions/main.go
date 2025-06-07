package main

import "fmt"

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

func main() {
	root := Sqrt(9)
	fmt.Println("Square root of 9 is:", root)
}
