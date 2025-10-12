package main

import "fmt"

// fn name args, return type
// typeが同じであれば省略可能
func add(x, y int) int {
	return x + y
}

// Multiple results
func swap (x, y string) (string, string) {
	return x, y
}

// Named return values
// Goでの戻り値となる変数に名前をつけることが可能
// 関数の最初で定義した変数名として扱われる
// この戻り値の名前は戻り値の意味を示す名前とすることで関数のドキュメントとして表現するようにする
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // neked return, x, yがreturnされる
}

func main() {
	fmt.Println(add(42, 13))
	a, b := swap("バルバッコア", "食べ行かない？")
	fmt.Println(a, b)

	fmt.Println(split(17)) // 7 10 x, yが出力される
}