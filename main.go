package main

import "fmt"

func main() {
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

}
