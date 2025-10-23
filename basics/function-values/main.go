package main

import (
	"fmt"
	"math"
)

// 関数も変数
// 他の変数のように関数を渡すことが可能
// 関数値(function value)は関数の引数にとることが可能、戻り値としても利用可能

func compute (fn func(float64, float64)float64) float64 {
	return fn(3, 4)
}


type BeefKind string

const (
	Sirloin    BeefKind = "sirloin"
	Ribeye     BeefKind = "ribeye"
	Tenderloin BeefKind = "tenderloin"
	Striploin  BeefKind = "striploin"
	Brisket    BeefKind = "brisket"
	Flank      BeefKind = "flank"
	Round      BeefKind = "round"
	Chuck      BeefKind = "chuck"
	ShortRib   BeefKind = "short rib"
	Rump       BeefKind = "rump"
)

func barbacoa(fn func(beefKind BeefKind) BeefKind) BeefKind {
	return fn("sirloin")
}

func main () {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	barbacoaStaff := func (beefKind BeefKind) BeefKind  {
		return beefKind
	}

	fmt.Println(barbacoa(barbacoaStaff))

}