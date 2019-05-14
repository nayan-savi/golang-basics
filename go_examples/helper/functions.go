package main

import (
	"fmt"
)

func test(x, y int) (int, int) {
	return y, x
}

func main() {
	x, _ := test(10, 20)
	fmt.Println(x)
}
