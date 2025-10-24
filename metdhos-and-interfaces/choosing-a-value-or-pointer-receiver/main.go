package main

import (
	"fmt"
	"math"
)

// Choosing a value or pointer recevier

// pointer receiverを使う2つの理由
// 1. methodがreceiverが指す先の変数を変更するため
// 2. methodの呼び出しごとに変数のコピーを避けるため


type Vertex struct {
	X, Y float64
}


func (v *Vertex) Scale(f float64) { 
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	sumX := v.X*v.X
	sumY := v.Y*v.Y
	sum := sumX + sumY
	fmt.Println("sum X and Y: ", sumX, sumY, "sum:", sum)
	sqrt := math.Sqrt(sum)
	fmt.Println("sqrt: ", sqrt)
	return sqrt
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}