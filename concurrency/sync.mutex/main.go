package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Mutex
// channelがgoroutine間で通信するための素晴らしい方法であることを見てきた
// しかし、通信が必要ない場合はどうか？コンフリクトを避けるために、
// 一度に1つのgoroutineだけが変数にアクセスできるようにしたい場合はどうか？

// このコンセプトは排他制御と呼ばれ、このデータ構造を指す一般的な名前はmutex
// Goの標準ライブラリは排他制御をsync.Mutexと次の二つのメソッドで提供する
// Lock
// Unlock
// Incメソッドにあるように、LockとUnlockで囲むことで排他制御で実行するコードを定義できる

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}