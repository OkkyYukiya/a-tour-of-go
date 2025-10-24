package main

import (
	"fmt"
	"math"
)

// Pointers and funcsions

// pointer argに取る、ScaleFunc関数はpointerを渡す必要がある
// var v Vertex
// ScaleFunc(v, 5)  // Compile error!
// ScaleFunc(&v, 5) // OK
// methodがpointer receiverである場合、引き出し時に変数、またはpointerのいずれかのreceiverとしてとることが可能

type Vertex struct {
	X, Y float64
}

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }


// func Scale(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func ScaleFunc (v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y *f
// }

// func main() {
// 	// v := Vertex{3, 4}
// 	// Scale(&v, 10)
// 	// fmt.Println(Abs(v))

//     v := Vertex{3,4}
// 	v.Scale(2)

// 	p := &Vertex{4,3}
// 	p.Scale(3)
// 	ScaleFunc(p, 8)
// 	fmt.Println(v, p)

// }


// Methods and pointer indirection (2)
// 逆にも見てみるぜ
// 変数の引数を取る関数は、特定の型の変数を取る必要がある

// 次に「変数の引数」を取る関数（＝参照で渡される）

// TypeScriptだと、オブジェクトを渡したときがこれに該当します。

// type Vertex = { x: number; y: number }

// function scale(v: Vertex, f: number) {
//   v.x *= f
//   v.y *= f
// }

// const v = { x: 3, y: 4 }
// scale(v, 10)

// console.log(v) // 👉 { x: 30, y: 40 } ← 変わってる！


// ここでの v は、オブジェクトの参照（つまり変数そのもの） が渡されます。
// だから scale() 内で v.x を変えると、元の v も変わる。

// 💡 これがGoでいう「変数の引数（＝ポインタ引数）」

// Goでは、オブジェクトではなく**明示的にアドレス（ポインタ）**を渡す必要があります。
// それが *Vertex のような形です。

// func Scale(v *Vertex, f float64) {
//     v.X = v.X * f
//     v.Y = v.Y * f
// }


// TSでは暗黙的に「参照で渡す」けど、
// Goは明示的に「この変数を直接いじっていいよ」と示すために * をつけます。

// 値レシーバー (value receiver)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 値引数（value argument）
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main () {
	v := Vertex{3,4}
	fmt.Println("v.Abs", v.Abs())
	fmt.Println("AbsFunc", AbsFunc(v))

	p := &Vertex{4,3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

}