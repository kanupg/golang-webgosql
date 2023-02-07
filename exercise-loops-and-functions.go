package main

import (
	"fmt"
	"math"
	"strconv"
)

const BASELINE_VALUE float64 = 0.01

func Sqrt(x float64) (float64, int) {
	z := 1.0
	prevVal := 1.0
	currentVal := 1.0
	count := 0
	for i := 0; i < 10; i++ {
		count++
		prevVal = z
		z -= (z*z - x) / (2 * z)
		currentVal = z
		if math.Abs(currentVal-prevVal) <= BASELINE_VALUE {
			//ごくわずかな変化の場合は処理を終了する
			break
		}
		// fmt.Println(z)
	}
	return z, count
}

func main() {
	for i := 0; i < 10; i++ {
		target := float64(i + 1)
		ans, count := Sqrt(target)
		fmt.Println(strconv.Itoa(i+1) + "を読み込み")
		fmt.Println(ans)
		fmt.Println(strconv.Itoa(count) + "回繰り返しました")

	}

}
