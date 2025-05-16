package main

import "fmt"

/**
* %d - 整数
* %s - 文字列
* %f - 浮動小数点数（例: %.2f で小数点以下2桁）
* %t - 真偽値（bool）
* %v - 値の既定フォーマット
* %+v - 構造体のフィールド名も含む
* %#v - Go構文での値の表現
* %T - 値の型
 */

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Printf("i is %d \n", i)
		sum += i
	}
	fmt.Println(sum)
}
