package main

import (
	"fmt"
	"sort"
	"stage1/readfile"
)

type TopK struct {
	data map[string]int
}

func NewTopK() *TopK {
	return &TopK{
		data: make(map[string]int),
	}
}
func (t *TopK) Add(word string) {
	t.data[word]++
}
func (t TopK) Top(n int) []string {
	sz := min(len(t.data), n)
	arr := make([]string, 0, len(t.data))
	for k := range t.data {
		arr = append(arr, k)
	}
	sort.Slice(arr, func(i, j int) bool {
		if t.data[arr[i]] == t.data[arr[j]] {
			return arr[i] < arr[j]
		}
		return t.data[arr[i]] > t.data[arr[j]]
	})
	return arr[:sz]
}
func (t *TopK) Count(word string) (int, bool) {
	value, ok := t.data[word]
	return value, ok
}
func main() {
	c := NewTopK()
	data, _ := readfile.ReadWords("./testdata/sample.txt")
	for _, s := range data {
		c.Add(s)
	}

	fmt.Println(c.Top(2))  // [go fun]   ("is" and "fun" tie at 1, "fun" < "is")
	fmt.Println(c.Top(10)) // [go fun is]  (fewer than 10 words exist; use your min trick)
}
