package main

import (
	"fmt"
	"math"
)

// Methods
// Goにはclassがnathing, だがしかし
// 型にmethodを定義することが可能
// methodは特別なレシーバ引数を関数に取る
// レシーバはfuncキーワードとmethod名の間に自身の引数リストで表現する


type Vertext struct {
	X, Y float64
}

// func (v Vertext) Abs() float64{
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// 通常関数ver
func abs (v Vertext) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


// Methods continued
// structの型だけでなく、任意の型(type)にもメソッドを宣言することが可能まじ？
// 例は, Absメソッドを持つ数値型のMyFloat型
// レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージにある必要がある
// 他のpackageに定義している型に対してレシーバを伴うメソッドは宣言できない


// tsの流儀、比較 type MyFloat = number;
// って思うやん？違うねん
// 実際は完全な別型。Goでは「同じ実体（数値）」でも「別の型」として扱われる。
// 以下に近いイメージ
// type Brand<T, Name> = T & { __brand: Name };
// type MyFloat = Brand<number, "MyFloat">;
type MyFloat float64

// (f MyFloat) レシーバ。this の代わり。
// f.Abs() こう呼び出す。

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main () {
	v := Vertext{3,4}
	fmt.Println(abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}