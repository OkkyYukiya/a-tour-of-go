package main

import "fmt"

// mapはkeyと値とを関連つける
// mapのzero値はnil
// nilマップはキーを持っておらず、キーを追加することができない
// make関数は指定された方のmapを初期化して、使用可能な状態で返す
// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex

// func main() {
// 	m = make(map[string]Vertex)

// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}

// 	fmt.Println(m["Bell Labs"])
// }

// var m map[string]Vertex
// m = make(map[string]Vertex)

// m["Bell Labs"] = Vertex{
//     40.68433, -74.39967,
// }

// イメージこんな感じ
// const m = {
//   "Bell Labs": { Lat: 40.68433, Long: -74.39967 }
// }

//Map literals
// map リテラルはstructリテラルに似ているが、キー(key)が必要である
// type Vertex struct {
// 	Lat, Long float64
// }

// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }

// func main() {
// 	fmt.Println(m)
// 	fmt.Println(m["Bell Labs"].Lat)
// }

// Map literals continued
// mapに渡すトップレベルの型が単純な型名である場合は
// リテラルの要素から推定できる、その型名を省略することが可能

// type Vertex struct {
// 	Lat, Long float64
// }

// var m = map[string]Vertex{
// 	"Bell Labs": {40.68433, -74.39967},
// 	"Google":    {37.42202, -122.08408},
// }

// func main() {
// 	fmt.Println(m)
// }

// Mutating Maps
// map mの操作いついて
// mへ要素(elem)の挿入や更新
// m[key] = elem

// 要素の取得
// elem = m[key]

// 要素の削除
// delete[m. key]

// キーに対する要素が存在するかどうかは、2つの目の値で確認します:
// elem, ok = m[key]
// elem, ok := m[key]



func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
