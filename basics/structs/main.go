package main

import "fmt"

// Structsはフィールドの集まり


type Vertex struct {
	X, Y int
}

// func main() {
// 	fmt.Println(Vertex{1, 2}) // {1 2}

// 	v:= Vertex{1, 2}
// 	v.X = 4
// 	fmt.Println(v.X) // 4
// 	// jsのobjectでミュータブルに変数を代入できるイメージだなこれわ
// 	// ただ型の厳格性は抑えとく必要がある, 以下はcompile error
// 	// untyped string constant) as int value in assignment
// 	// v.X = "バルバッコア"
// }

// structのフィールドは、structのポインタを通してアクセスすることが可能
// フィールドxを持つstructのポインタpがある場合
// フィールドxにアクセスするには(*p).xのように書くことが可能
// しかし表記法は面倒なためGoでは代わりにp.Xと書くことが可能
// func main() {
// 	v := Vertex{1, 2}
// 	p := &v
	
// 	fmt.Println(&v) 
// 	fmt.Printf("%p\n", p)
// 	fmt.Printf("%p\n", &v)
// 	// &{1 2} //Printlnは自動的にdereferenceしてlogを出す、&がprefixについてpinterであることを示す
// 	// 0x140000b4010
// 	// 0x140000b4010
// 	p.X = 1e9
// 	fmt.Println(v)
// }


// Struct Literals
// struct リテラルはフィールドの値を列挙することで新しいstructの初期値の割り当てを示す
// Name: 構文を使って、フィールドの一部だけを列挙することが可能
// &を頭につけると新しく割り当てられたstructへのpointerを戻す


var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y: 9 is implicit
	v3 = Vertex{} // X:0, Y:0
	p = &Vertex{1, 2} // has type *Vertex
)
func main () {
  fmt.Println(v1, v2, v3, p)
  // {1 2} {1 0} {0 0} &{1 2}
}