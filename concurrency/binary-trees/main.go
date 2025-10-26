package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// 1. Walkの実装
// walkRecursive はtreeを「中順（in-order）」で再帰的に探索
// 見つけた値をチャネル ch に送信する関数
func walkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		// 葉ノードの先に到達したら終了させる
		return 
	}
	
	// 中順探索 (Left->Node->Right)
	walkRecursive(t.Left, ch)  // 1. 左の部分木を探索
	ch <- t.Value              // 2. 現在のノードの値を送信
	walkRecursive(t.Right, ch) // 3. 右の部分木を探索
}

// Walk はtree tを探索し、見つかったすべての値をチャネル ch に送信するでやんす
// すべての値を送信したらチャネルをclose
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	// 探索がすべて完了したら（walkRecursiveが終了したら　受信側（range）chを閉じる
	close(ch)
}

// 3. Same 関数の実装
// Same は2つの木 t1 と t2 が同じ値を含んでいるかどうかを判断します。
func Same(t1, t2 *tree.Tree) bool {
	// 2つの木に対応する2つのチャネルを作成
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 2つのゴルーチンを起動し、それぞれの木を並行して探索
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// 無限ループで両方のチャネルを監視
	for {
		// チャネルから値を受信
		// v  <- 値
		// ok <- チャネルが開いていれば true、閉じていれば false
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// --- 状態チェック ---

		// 1. 両方のチャネルが閉じ、かつ値も同じだった
		//    (ok1 も ok2 も false)
		if !ok1 && !ok2 {
			return true // 2つの木は同じ
		}

		// 2. どちらか一方のチャネルだけが閉じた
		//    (木のノード数が異なる)
		if ok1 != ok2 {
			return false // 2つの木は異なる
		}

		// 3. 両方のチャネルから値を取得できたが、値が異なる
		if v1 != v2 {
			return false // 2つの木は異なる
		}

		// 4. 値が同じ -> ループを続けて次の値を比較
	}
}

// ------------------------------------------------------------------
// 2 & 4. テストの実行
// ------------------------------------------------------------------

func main() {
	// 2. Walk 関数のテスト
	fmt.Println("--- 1. Walk 関数のテスト (tree.New(1)) ---")
	ch := make(chan int)
	go Walk(tree.New(1), ch) // run Walk by goroutine

	// range ch はチャネル ch が close されるまで値を受信し続ける
	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println("\n(期待値: 1 2 3 4 5 6 7 8 9 10)")

	// 4. Same 関数のテスト
	fmt.Println("\n--- 2. Same 関数のテスト ---")

	// テスト 1: tree.New(1) と tree.New(1)
	// 構造はランダムで異なる可能性があるが、
	// 保持する値は (1, 2, ..., 10) で同じはずやんね
	fmt.Printf("Same(tree.New(1), tree.New(1)): %v\n", Same(tree.New(1), tree.New(1)))

	// テスト 2: tree.New(1) と tree.New(2)
	// tree.New(1) は (1, 2, ..., 10)
	// tree.New(2) は (2, 4, ..., 20)
	// 保持する値が異なるため、false になるはずバルバッコア
	fmt.Printf("Same(tree.New(1), tree.New(2)): %v\n", Same(tree.New(1), tree.New(2)))
}