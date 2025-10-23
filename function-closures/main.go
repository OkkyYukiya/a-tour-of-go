package main

import "fmt"

// Function closures
// Goの関数はクロージャ、それ自身の外部から変数を参照する関数値
// この関数は参照された変数へアクセスして帰ることができる、その関数は変数へbindされている



func adder () func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		fmt.Println("sum: ", sum)
		return sum
	}
}


func  fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main () {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(i, "-----")
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}