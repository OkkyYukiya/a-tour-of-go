package main

import "fmt"

// defer statementはdeferへ渡した関数の実行を
// 呼び出し元の関数の終わり(returnする)まで遅延させる
// deferへ渡した関数の引数はすぐに評価されるが、その関数自体は呼び出し元の関数がreturnされるまで実行されない
func main () {
    defer fmt.Println("world")
	fmt.Println("hello")
}

// deferを使うケース リソースのクリーンアップ（最も重要）, ロックの解放、パニックからの回復など
// func complex() error {
//     file.Open()
//     defer file.Close()  // 一度だけ書けばOK
    
//     if err1 != nil {
//         return err1  // 自動的にClose()される
//     }
    
//     if err2 != nil {
//         return err2  // ここでも自動的に
//     }
    
//     return nil  // ここでも自動的に
// }