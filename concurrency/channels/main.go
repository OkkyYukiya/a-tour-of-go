package main

import "fmt"

// Channels
// チャネル型は、チャネルオペレータの <- を用いて値の送受信ができる通り道
// ch <- v    // v をチャネル ch へ送信する
// v := <-ch  // ch から受信した変数を v へ割り当てる
// (データは、矢印の方向に流れる)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := [6]int{7, 2, 8, -9, 4, 0}

	// マップとスライスのように、チャネルは使う前に以下のように生成します
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
