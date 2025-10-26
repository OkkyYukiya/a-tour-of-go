package main

import "fmt"

// Range and Close

// 送り手はこれ以上の送信する値がないことを示すため
// チャネルをcloseすることが可能
// 受けては、受信の式に2つ目のパラメータを割り当てることでそのチャネルがcloseされているかを確認できる

// v, ok := <-ch

// 受信する値がない、チャネルが閉じているなら、okの変数はfalseになる

// loopのfor i := range cはチャネルが閉じられるまでチャネルから値を繰り返し受信し続ける

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// あとで確認する
func main() {
	c := make(chan int, 10)
	fmt.Println(c)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}