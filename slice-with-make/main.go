package main

import (
	"fmt"
)

// sliceは組み込みのmake関数を利用することで作成可能
// func main () {
// 	// make関数はzero化された配列を割り当て、その配列を指すスライスを返す
// 	// makeの3番目の引数に、スライスのcapを指定可能
//     a := make([]int, 5) // 動的サイズの配列を作成する方法
// 	printSlice("a", a)

// 	b := make([]int, 0, 5)
// 	printSlice("b", b)

// 	c:=b[:2]
// 	printSlice("c", c)

// 	d:=c[2:5]
// 	printSlice("d", d)

// }

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }

// slices of slices

// func main () {
//     // Create a tic-tac-tae board

// 	board := [][]string{
// 		[]string{"_","_","_"},
// 		[]string{"_","_","_"},
// 		[]string{"_","_","_"},
// 	}

// 	fmt.Println(board, len(board[0]), cap(board))

// 	board[0][0] = "X"
// 	board[2][2] = "O"
// 	board[1][2] = "X"
// 	board[1][0] = "O"
// 	board[0][2] = "X"

// 	for i := 0; i < len(board); i++ {
// 		fmt.Printf("%s\n", strings.Join(board[i], " "))
// 	}

// }

// Appending to a Slice
// スライスへ新しい要素を追加するにはGoの組み込みのappendを使う

// func append(s []T, vs ...T)[]T
// sは追加元となるT型のスライス、残りのvsは追加するT型の変数群
// appendの戻り値は追加元のsliceと追加する変数群を合わせたslice
// もし、元の配列sが変数群を追加する際に容量が小さい場合より大きサイズの配列を割り当て直す


func main () {
	var s []int // len=0 cap=0 []
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s) // len=1 cap=1 [0]

	// the slice grows as needed.

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2,3,4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
