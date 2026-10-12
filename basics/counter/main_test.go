package main

import (
	"slices"
	"strings"
	"testing"
)

func TestTop2(t *testing.T) {
	c := NewTopK()
	data := "go is go fun"
	for s := range strings.FieldsSeq(data) {
		c.Add(s)
	}
	got := c.Top(2)
	need := []string{"go", "fun"}
	if slices.Compare(got, need) != 0 {
		t.Errorf("Top(2) = %v, want %v", got, need)
	}
}
func TestTop10(t *testing.T) {
	c := NewTopK()
	data := "go is go fun"
	for s := range strings.FieldsSeq(data) {
		c.Add(s)
	}
	got := c.Top(10)
	need := []string{"go", "fun", "is"}
	if slices.Compare(got, need) != 0 {
		t.Errorf("Top(10) = %v, want %v", got, need)
	}
}
func TestTopEmpty(t *testing.T) {
	c := NewTopK()
	expected := 0
	got := len(c.Top(5))

	if got != expected {
		t.Errorf("expected len %d but instead got %d", expected, got)
	}
}
func TestTopTie(t *testing.T) {
	c := NewTopK()
	data := "b a c"
	for s := range strings.FieldsSeq(data) {
		c.Add(s)
	}
	got := c.Top(10)
	expected := []string{"a", "b", "c"}
	if slices.Compare(got, expected) != 0 {
		t.Errorf("Top(10) = %v, want %v", got, expected)
	}
}
func TestTopCount(t *testing.T) {
	c := NewTopK()
	data := "go is go fun"
	for s := range strings.FieldsSeq(data) {
		c.Add(s)
	}
	value, ok := c.Count("go")
	expectedCnt, expectedOk := 2, true
	if value != expectedCnt || ok != expectedOk {
		t.Errorf("got (%d, %v), want (%d, %v)", value, ok, expectedCnt, expectedOk)
	}
	value, ok = c.Count("rust")
	expectedCnt, expectedOk = 0, false
	if value != expectedCnt || ok != expectedOk {
		t.Errorf("got (%d, %v), want (%d, %v)", value, ok, expectedCnt, expectedOk)
	}
}
