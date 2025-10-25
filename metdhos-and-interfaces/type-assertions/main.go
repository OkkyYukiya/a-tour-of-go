package main

import "fmt"

// Type assertions
// 型アサーションはinterfaeの型の基になる具体的な値を利用する手段を提供する

// ```
// t := i.(T)
// ```

func main() {
	// i が T を保持していれば、 t は基になる値になり、 ok は真(true)になります。
	// そうでなければ、 ok は偽(false)になり、 t は型 T のゼロ値になり panic は起きません。
	var i interface{} = "hello"
	// var i interface{}
	s := i.(string)
	// var s string
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)
	// hello true

	f, ok := i.(float64)
	fmt.Println(f, ok)
	// 0 false, zero値
}