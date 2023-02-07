package main

import (
	"fmt"
	"math"
)

func main() {
	flg := true
	fmt.Println(checkFlg(flg))

	fmt.Println(
		pow(3, 2, 10),
	)
}

func checkFlg(target bool) string {
	//基本的なif文
	if target {
		return "真です"
	} else {
		return "偽です"
	}

}

func pow(x, n, lim float64) float64 {

	//条件式の直前に簡単なステートメントを記述できる
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// ここでは変数vは使えない
	return lim
}
