package main

import (
	"fmt"
	"math"
)

// Pointer Recivers
// 値レシーバー: 値を読み取りたい、コピーしたいため
// pointer reveivers: structを直接いじる、同じメモリ上で動作する
type Vertex struct {
	X, Y float64
}


func (v Vertex) Abs() float64 {
	v1 := v.X*v.X
	v2 := v.Y*v.Y
	println(v1, v2)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale (f float64) {
	v.X = v.X *f
	v.Y = v.Y *f
}


func main () {
	v := Vertex{3,4}
	fmt.Printf("%v, %v,\n", v.X, v.Y)
	v.Scale(10)// X, Yが10倍された値になる
	fmt.Printf("%v", v)
	fmt.Println(v.Abs())
}