package main

import (
	"fmt"
)

func main() {
	// 10回の繰り返し
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 配列を繰り返し
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, value := range a {
		fmt.Println(i, ":", value)
	}

	b := 0
	for b < 10 {
		fmt.Println(b)
		b++
	}
}
