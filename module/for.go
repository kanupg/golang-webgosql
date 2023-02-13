package main

import "fmt"

func main() {
	sum := 0
	//丸括弧がいらない
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//初期化と後処理ステートメントとセミコロン省略可能
	sum = 1
	for sum < 10 {
		sum += 1
	}
	fmt.Println(sum)

	//無限ループはかんたん
	for {
		fmt.Println("無限")
		break
	}

}
