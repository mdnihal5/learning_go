package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Word Counter :")
	text := "go is fun and go is fast and go"
	counter := make(map[string]int)
	for v := range strings.FieldsSeq(text) {
		counter[v]++
	}
	fmt.Println(len(counter))
	slice := make([]string, 0, len(counter))
	for k := range counter {
		slice = append(slice, k)
	}
	sort.Slice(slice, func(i, j int) bool {
		if counter[slice[i]] == counter[slice[j]] {
			return slice[i] < slice[j]
		}
		return counter[slice[i]] > counter[slice[j]]
	})
	for _, v := range slice {
		fmt.Println(v, counter[v])
	}

}
