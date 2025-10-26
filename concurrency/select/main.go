package main

import (
	"fmt"
	"time"
)

// Select
// select statementはgoroutineを複数の通信操作で待たせる
// selectは複数あるcaseのいずれかが準備できるようになるまでブロックし
// 準備ができたcaseを実行する
// もし複数のcaseの準備ができている場合、caseはランダムに選択される

// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		select {
// 		case c <- x:
// 			x, y = y, x+y
// 		case <- quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c)
// 		}
// 		quit <- 0
// 	}()

// 	fibonacci(c, quit)
// }

// Default Selection
// どの case も準備ができていないのであれば
// select の中の default が実行される
// ブロックせずに送受信するならdefault の case を使う

func main(){
	tick := time.Tick(100 * time.Microsecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Microsecond)
		}
	}
}