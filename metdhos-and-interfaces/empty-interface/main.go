package main

import "fmt"

// The empty interface
// 0個のメソッドを指定されたinterface型は、空のinterfaceと呼ばれる
// interface{}

// 空のinterfaceは任意の型の値を保持できる
// (全ての型は、少なくとも0個のメソッドを実装している)

// 型のinterfaceは、未知の方の値を扱うコードで使用される
// 例えば、fmt.Printはinterface{}型の任意の数の引数を受け取る

func main() {
	var i interface{}
	describe(i)
	// (<nil>, <nil>)

	i = 42
	describe(i)
	// (42, int)

	i = "hello"
	describe(i)
	// (hello, string)
}


func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}