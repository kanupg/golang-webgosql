package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	for i := 1; i <= 10; i++ {
		result, cnt := Sqrt(float64(i))
		fmt.Println(strconv.Itoa(i) + "を投入します")
		fmt.Println(result)

		if cnt < 3 {
			fmt.Println("3回以内")
		}
		fmt.Println("")

	}

}

func Sqrt(x float64) (float64, int) {
	z := 1.0
	prevVal := 1.0
	cnt := 0

	for i := 0; i < 10; i++ {
		cnt++
		prevVal = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		fmt.Println(strconv.Itoa(cnt) + "回目の計算する")
		if math.Abs(z-prevVal) < 0.000000001 {
			break
		}
	}

	return z, cnt

}
