package main

import (
	"fmt"
)

// interface values
// 下記のように、interfaceの値は値と具体的な型のタプルのように考えることが可能
// 改めてtableは型が固定、順序がある、普遍
// (value type)
// interfaceの値は、特定の基底になる具体的な型の値を保持する
// interfaceの値のmethodを呼び出すと、その基底型の同じ名前のメソッドが実行される

type I interface {
	M()
}

type T struct {
	S string
}

// func (t *T) M() {
// 	fmt.Println(t.S)
// }

// type F float64

// func (f F) M() {
// 	fmt.Println(f)
// }

// func main() {
// 	var i I
// 	// fmt.Println(i.M()) // この時点で実行するとcompile error, valueがない

// 	i = &T{S: "Hello"}
// 	describe(i)
// 	i.M()
// 	i = F(math.Pi)
// 	describe(i)
// 	i.M()
// }


func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}


// Interface values with nil underlying values 
// interface自体の中になる具体的な値がnilの場合、メソッドはnilをレシーバーとして呼び出す
// いくつかの言語でこれはnull pointer例外を引き起こすが、Goではnilをレシーバーとして呼び出されても適切に処理するうメソッドを記述するのが一般的
// 具体的な値としてnilを保持するinterfaceの値のそれ自体は非nilであることに注意する

// func (t *T) M() {
// 	if t == nil {
// 		fmt.Println("<nil>")
// 		return 
// 	}
// 	fmt.Println(t.S)
// }

// func main () {
// 	var i I
// 	var t *T
// 	i = t
// 	describe(i)
// 	i.M()

// 	i = &T{"hello"}
// 	describe(i)
// 	i.M()
// 	// 	(<nil>, *main.T)
// 	// <nil>
// 	// (&{hello}, *main.T)
// 	// hello
// }


// Nil interface values
// nil intafeceの値は、値も具体的な型も保持しない
// 呼び出す具体的なメソッドを示す型がinterfaceのタプル内に存在しないため
// nil interfaceのメソッドを呼び出すと、ランタイムエラーになる

func main () {
	var i I
	describe(i)
	// i.M()
	//(<nil>, <nil>)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x101006a38]

	// goroutine 1 [running]:
	// main.main()
}