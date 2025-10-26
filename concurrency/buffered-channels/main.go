package main

import "fmt"

// Beffered Channles

// チャネルは、bufferとして使える
// チャネルを初期化するには、makeの2つ目の引数にbufferの長さを与える

// ch := make(chan int, 100)

// bufferが詰まった時は、チャネルへの送信をブロックする
// bufferが空の時にはチャネルの受信をブロックする

func main() {
	ch := make(chan int, 1)
	// fatal error: all goroutines are asleep - deadlock!
	// ch := make(chan int, 100)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}