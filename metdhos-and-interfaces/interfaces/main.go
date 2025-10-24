package main

// Interfeces

// interface型はメソッドのシグニチャの集まりで定義する

// BYGPT
// Goにおける インターフェースのシグネチャ（signature） は、
// 「そのメソッドがどんな引数を取り、どんな値を返すか」という“型上の約束” のことです。
// つまり、「インターフェースに定義されたメソッドの定義部分（関数の顔）」＝シグネチャ。

import (
	"fmt"
)

type Abser interface {
	Abs() float64
}

// func main() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}

// 	a = f  // a MyFloat implements Abser
// 	a = &v // a *Vertex implements Abser

// 	// In the following line, v is a Vertex (not *Vertex)
// 	// and does NOT implement Abser.
// 	a = &v

// 	fmt.Println(a.Abs())
// }

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }


// Interfaces are implemened implicitly
// 型にメソッドを実装していくことによって、
// interfeceを実装する(満たす)
// 暗黙のinterface, interfaceの定義をその実装から切り離す、interfaceの実装は事前の取り決めなしにpackageに現れることがある

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I, 
// but we don't neet to explicitly declare that it dose so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Barubacoa"}
	i.M()
}

