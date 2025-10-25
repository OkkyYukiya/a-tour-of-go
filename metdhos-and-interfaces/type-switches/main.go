package main

import "fmt"

// Type switches

// 型switchはいくつかの型アサーションを直列に使用できる構造
// 型switchは通常のswitchと似ているが、型swtichのcaseは型を指定する(値でない)
// それらの値は指定されたinterfaceの値が保持する値の型と比較される
// ```
// switch v := i.(type) {
// case T:
//     // here v has type T
// case S:
//     // here v has type S
// default:
//     // no match; here v has the same type as i
// }
// ```
// 型switchの宣言は型アサーションi.(T)と同じ構文を持つが、特定の型Tはキーワードtypeに置き換えられる

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("q %v is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}


func main() {
	do(21)
	do("barbacoa")
	do(true)
}