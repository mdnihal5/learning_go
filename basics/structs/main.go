package main

import "fmt"

type Count struct{ n int }

func (c Count) IncV() {
	c.n++
	fmt.Println(c.n)
}
func (c *Count) IncP() {
	c.n++
}
func main() {

}
