package main

/** Exported Names */
// import "fmt"
// import "math"
// equal
// ↓ import statements is grouped
import (
	"fmt"
	"math"
	"math/rand"
)

/** Functions */
// argsは x typeとする
// x int, y int = x, y intと同等
func add(x, y int) int {
	return x + y
}

//  Typescript
//  function add (x: int, y: int) {
//	    return x + y
//  }

/** Multiple results */
func swap(x, y string) (string, string) {
	return x, y
}

/** Named return values */
// naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

/** Variable declaration */
var c, python, java bool

// variables with initializers
// var i, j int = 1, 2

// short variable declaration
// := ショート代入, 暗黙的な型宣言が可能
// k := 3 関数の中でしか使えない

/** Basic Types */
// 		bool
// 		string
// 		int  int8  int16  int32  int64
// 		uint uint8 uint16 uint32 uint64 uintptr
//         - int, uint, uintptr 型は32-bitのシステムでは32 bitで、64-bitのシステムでは64bitになる
// 				 - サイズ、符号なし( unsigned )整数の型を使うための特別な理由がない限り、整数の変数が必要な場合はintを使うのがいい
//		byte // uint8 の別名
// 		rune // int32 の別名 runeとは古代文字らしい、　Goでは文字そのものを表すためにruneという言葉を使う
// 		     // Unicode のコードポイントを表す
// 		float32 float64
// 		complex64 complex128

/** Zero values */
// 変数に初期値を与えない場合、zero値が与えられる
// - 数値型(int,floatなど): 0
// - bool型: false
// -- string: ""

/** Type conversions */
// 型変換
// typescriptのNumber()と同じイメージ
// var _i int = 42
// var _f float64 = float64(_i)
// var _u uint = uint(_i)

/** Type Interface */
// 明示的な型を指定せずに変数を宣言するケースでは右辺の変数から型推論される

/** Constants */
// 定数
// character, string, bool, numericのみで利用可能
// :=は使えないよ
const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println((math.Pi))

	fmt.Println(add(42, 13))

	fmt.Println(swap("hello", "world"))

	fmt.Println(split(17)) // x=7, y=10

	// k := 3
	// fmt.Println(i, j, k, c, python, java)

	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)

	const world = "世界"
	fmt.Println("Hello", world)
	const Truth = true
	fmt.Println("Go rules?: ", Truth)

	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}
