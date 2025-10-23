package main

import "fmt"

// [n]T型は型Tの個の変数の配列を表す
// 以下は、int10個の配列を宣言
// var a[10]int
// 配列の長さは、型の一部分、配列のサイズは変更不可
// func main() {
// 	var a[2]string
// 	fmt.Println(a)
// 	fmt.Println(len(a)) // 2とoutputされる
// 	fmt.Printf("type: %T\n", a[0])
// 	if a[0] == "" {
// 		fmt.Println("empty string")
// 	}
// 	a[0] = "hoge"
// 	a[1] = "huga"
// 	fmt.Println(a)

// 	primes := [6]int{2,3,5,7,11,13}
// 	fmt.Println(primes)
// }

// Slices
// 配列は固定長、sliceは可変長
// より柔軟な配列、スライスは配列よりも一般的
// 型 []TはTのスライスを表す
// 転んで区切られた二つのインデックスlowとhighの境界を指定することによってスライスが形成されます
// a[low : high]
// これは最初の要素は含むが、最後の要素はのぞいた半開区間を選択する
// 次の式はaの要素の内、1~3を含むスライスを作成する
// a[1:4]

// func main() {
//     primes := [6]int{2,3,5,7,11,13}
// 	var s []int = primes[1:4]
// 	fmt.Println(s) // [3 5 7]
// }

// Slices are like references to arrays
// スライスは配列への参照のようなもの
// スライスはどんなデータも格納していない、単に元の配列の部分列を指し示す
// スライスの要素を変更すると、その元となる配列の対応する要素が変更される
// 同じ元となる配列を共有している他のスライスはそれらの変更が反映される

// func main () {
// 	babacoa := [4]string {
// 		"サーロイン",
// 		"リブロース",
// 		"イチボ",
// 		"チーズ",
// 	}
// 	fmt.Println(babacoa)
// 	// [サーロイン リブロース イチボ チーズ]

// 	a := babacoa[0:2]
// 	b := babacoa[1:3]
// 	fmt.Println(a, b)
// 	// [サーロイン リブロース] [リブロース イチボ]

// 	b[0] = "カレーライス"
// 	fmt.Println(a, b)
// 	fmt.Println(babacoa)
// 	// [サーロイン カレーライス] [カレーライス イチボ]
//     // [サーロイン カレーライス イチボ チーズ]
//     //　元の配列babacoa自体も上書きされる
// }

// Slice literals
// スライスのリテラルは長さのない配列のリテラルのようなもの
// これは配列リテラル
// [3]bool{true, true, true}
// これは上記と同様の配列を作成し、それを参照するスライスを作成する
// []bool{true,true,false}

func main () {
    q := []int{2,3,5,7,11,13}
	fmt.Println(q)

	r := []bool{true,false,true,true,false,true}
	fmt.Println(r)


	s := []struct{
		i int
		b bool
 	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Println(i, "times", s[i].i)
	}
}