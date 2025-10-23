package main

import "fmt"

// Pointers
// pointerは値のメモリアドレスを指す
// pointerは値のメモリアドレスを指す
// pointerは値のメモリアドレスを指す

// var Tのpointerは*T型でゼロ値はnil
// var p *int

// &オペレータはそのオペランド(operand)へのポインタを引き出す
// *オペレータは、pointerの指す先の変数を示す

// &(アドレス演算子)はその値が格納されているメモリのアドレスを参照する
// *(dereference演算子)はpointerが指すアドレスの実際の値を参照する
func main () {
	// i := 42
	// p := &i
	// fmt.Println(&i) // adressをouput 0Xxxxxxxx
	// fmt.Println(*p) // 値をoutput 42
	// *p = 21

	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
