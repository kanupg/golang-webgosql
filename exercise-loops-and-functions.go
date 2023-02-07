package main

import (
	"fmt"
	"strconv"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		//fmt.Println(z)
	}
	return z
}

func main() {
	for i := 0; i < 10; i++ {
		target := float64(i + 1)
		fmt.Println(strconv.Itoa(i+1) + "を読み込み")
		fmt.Println(Sqrt(target))
	}

}
