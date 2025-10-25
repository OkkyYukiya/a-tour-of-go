package main

import (
	"fmt"
	"time"
)

// Goのプログラムは、エラーの状態をerror値で表現する
// error型のfmt.Stringerに似た組み込みのinterfaceである
// type error inrerface { Error() string }

type MyError struct {
	When time.Time
	What string
}


func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}


func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// func main() {
// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 		// at 2025-10-25 15:24:10.426502 +0900 JST m=+0.000068167, it didn't work
// 	}
// 	// nilのerrorは成功したことを示す、nilでない場合、それはつまり失敗を表す、バルバッコアに行けなかった時のような気持ち
// }

type ErrNegativeSqrt float64

// error インタフェースを満たす
func (e ErrNegativeSqrt) Error() string {
	// fmt.Sprint(e) を直接呼ぶと無限ループになる（なぜなら fmt がまた e.Error() を呼ぼうとするため）
	// float64(e) にキャストして数値として出力する
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// errorを返すようにする
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		// ニュートン法を用いて平方根を求める 以下引用
		// コンピュータは通常ループを使って x の平方根を計算します。
		// いくつかの z を推測することから始めて、
		// z² がどれほど x に近づいているかに応じて z を調整できます。

		result := z - (z*z-x)/(2*z)

		// 収束条件として前回の値と一致したら解が出たと見て処理を終了させる楽しいなこれ
		if result == z {
			return z, nil
		}
		z = result
		fmt.Println("i: ", i, "z: ", z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
