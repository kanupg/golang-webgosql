package main

import (
	"fmt"
)

// 定数 頭文字大文字で他のパッケージでも使える
const Pi = 3.14

const (
	URL      = "http://xxx.exmanple.com"
	SiteName = "titleeeee"
)

func main() {

	/*
		fmt.Println("Hello World")

		//明示的な定義
		//var 変数名 型 = 値
		var i int = 100
		fmt.Println(i)

		//同じ型をまとめて定義
		var t, f bool = true, false
		fmt.Println(t, f)

		//異なる型をまとめて定義
		var (
			i2 int    = 200
			s2 string = "Golang"
		)
		fmt.Println(i2, s2)

		//暗黙的な定義（※関数内でしか定義できない）
		//変数名 := 値
		i4 := 400
		fmt.Println(i4)

		//配列
		var arr1 [3]int
		fmt.Println(arr1)

		var arr2 [3]string = [3]string{"aaa", "bbb", "ccc"}
		fmt.Println(arr2)

		arr3 := [3]string{"a", "b", "c"}
		fmt.Println(arr3)

		arr4 := [...]string{"ddd", "eee"}
		fmt.Println(arr4)

		//取り出し
		fmt.Println(arr2[1])

		//len
		fmt.Println(len(arr4))

		//文字列→数値
		i, _ = strconv.Atoi("1000")
		fmt.Println(i)

		str1 := strconv.Itoa(i)
		fmt.Println(str1)

	*/

	//関数
	funcTest1, funcTest2 := Plus(1, 2)
	fmt.Println(funcTest1, funcTest2)

	// func funcTest(x,y int) int {

	// }

} //main

func Plus(a, b int) (int, int) {
	ans1 := a + b
	ans2 := a + b + 1
	return ans1, ans2
}
