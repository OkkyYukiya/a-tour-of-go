package main

import "fmt"

func main () {
	sum := 1
    // for (let i = 0, i < 10, i ++) {}と一緒やで
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("i is %v\n", i)
	// 	sum += i
	// }
   
	//;省略可能
	for sum < 1000 {
		sum += sum
	}	
	fmt.Println(sum)
}