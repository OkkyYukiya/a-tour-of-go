package main

import "fmt"

// Stringers

// 最もよく利用されるinterfaceの一つに、fmt packageに定義されているStringersがある
// Stringers interfaceはstringとして表現することができる型です
// fmt packageでは、変数を文字列で出力するためにこのinterfaceがあることを確認する

// type Person struct {
// 	Name string
// 	Age int
// }

// func (p *Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }

// func main() {
// 	a := &Person{"Authur Dent", 42}
// 	z := &Person{"Zaphod Barbacoa", 9001}

// 	fmt.Println(a, z)
// 	// Authur Dent (42 years) Zaphod Barbacoa (9001 years)
// }

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
