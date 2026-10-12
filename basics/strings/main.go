package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "go is   fun and fast"
	slice := strings.Fields(s)
	for _, sub := range slice {
		fmt.Println(sub)
	}
	fmt.Println(len(slice))
	slice = strings.Split(s, " ")
	fmt.Println()
	for _, sub := range slice {
		fmt.Println(sub)
	}
	fmt.Println(len(slice))
}
