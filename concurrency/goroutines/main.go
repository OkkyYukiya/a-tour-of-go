package main

import (
	"fmt"
	"time"
)

// Goroutines
// goroutineはGoのruntimeに管理される軽量なスレッド

// go f(x, y, z) と書けば、新しいgoroutineが実行されます

// f(x, y, z)
// f , x , y , z の評価は、実行元(current)のgoroutineで実行され、 f の実行は、新しいgoroutineで実行されるまじ？
// goroutineは同じアドレス空間で実行されるため、共有メモリへのアクセスは必ず同期する必要がある
// sync packageは同期する際に役にたつ方法を提供していますが、別の方法があるため、それほど必要ない

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}


func main() {
	go say("world")
	say("hello")
	// hello
	// world
	// world
	// hello
	// hello
	// world
	// world
	// hello
	// hello
}