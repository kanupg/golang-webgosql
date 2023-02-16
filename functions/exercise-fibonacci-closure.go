package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	idxN, idxNPlus1 := 0, 1
	return func() int {
		ans := idxN
		idxN = idxNPlus1
		idxNPlus1 = idxN + ans
		return ans
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
