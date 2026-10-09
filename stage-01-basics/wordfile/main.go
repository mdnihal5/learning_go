package main

import (
	"fmt"
	"sort"
	"stage1/readfile"
)

func main() {
	data, err := readfile.ReadWords("./testdata/short.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	counter := make(map[string]int)
	for _, k := range data {
		counter[k]++
	}
	unq := make([]string, 0, len(counter))
	for k := range counter {
		unq = append(unq, k)
	}
	sort.Slice(unq, func(i, j int) bool {
		if counter[unq[i]] == counter[unq[j]] {
			return unq[i] < unq[j]
		}
		return counter[unq[i]] > counter[unq[j]]
	})

	sz := min(len(unq), 5)
	unq = unq[:sz]
	fmt.Println(unq) // [go fun is]

	data2, err := readfile.ReadWords("nofile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data2)
}
