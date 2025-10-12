package main

import (
	"fmt"
	"math/cmplx"
)

// BasicTypes
// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // uint8 の別名

// rune // int32 の別名
//      // Unicode のコードポイントを表す

// float32 float64

// complex64 complex128

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {return x*10 + 1}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main () {
	fmt.Printf("Type %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value %v\n", z, z)
	// Type bool Value false
	// Type uint64 Value 18446744073709551615
	// Type complex128 Value (2+3i)

	// 以下0値
	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// 0 0 false "" 

	// Type conversions 型変換
	// var v, type Tがあった場合、T(v)は、変数vをT型へ変換する
	// tsでいうNumber()的なやつだなこれは

	var i int = 42
	fmt.Printf("Type %T Value %v\n", i, i)
	var f float64 = float64(i)
	fmt.Printf("Type %T Value %v\n", f, f)
	u := uint(f)
	fmt.Printf("Type %T Value %v\n", u, u)
	// Type int Value 42
	// Type float64 Value 42
	// Type uint Value 42

	// Type inference
	// 型推論的なやつ、代入する右辺を元に型推論が実行されるぜ
	// :=, var = のいずれか

	v := 42
	fmt.Printf("Type %T", v) // int

	// Constants
	// 定数はconstキーワードを利用して宣言
	// 定数は文字(character)、文字列(string)、bool, numericのみで利用可能
	// :=は使えない

	// const Heaven = "バルバッコア"
	// fmt.Println("Hello", Heaven)

	// Numeric Constants
	// 数値の定数は、高精度な値(values)です
	// 型のない定数はその状況によって必要な型をとることになる
	fmt.Println(needInt((Small)))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

