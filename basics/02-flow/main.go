package main

import (
	"fmt"
	"math"
)

// fmtのvalue置換
// - %d - 整数
// - %s - 文字列
// - %f - 浮動小数点数（例: %.2f で小数点以下2桁）
// - %t - 真偽値（bool）
// - %v - 値の既定フォーマット
// - %+v - 構造体のフィールド名も含む
// - %#v - Go構文での値の表現
// - %T - 値の型
func PrintFormattedLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Printf("i is %d \n", i)
		sum += i
	}
	fmt.Println(sum)
}

// for ; sum < 1000;
// 初期化(i := 0), 後処理(i++)の記述は任意
// セミコロンを省略することも可能 for sum < 1000
// while文はない、Goではforを用いる
func InitializeLoop() {
	sum := 1
	for sum < 1000 {
		// for sum < 1000 省略version
		sum += sum
	}

	// 無限loop
	// for {
	// }
}

// if statement
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// v := math.Pow(x, n)
// if statementは条件の前に評価するためのstatementを記述可能
// scopeはif内のみ
func pow(x, n, lim float64) float64 {
	fmt.Println(math.Pow(x, n))
	if v := math.Pow(x, n); v < lim {
		fmt.Println(v)
		return v
	}

	return lim
}
func main() {
	// PrintFormattedLoop()
	// fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
