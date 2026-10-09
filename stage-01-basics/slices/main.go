package main

import (
	"fmt"
)

func main() {
	b := []int{1, 2, 3}
	a := make([]int, len(b))
	copy(a, b)
	a[0] = 99
	fmt.Println(a, b)
	fmt.Println(a, b)
}
