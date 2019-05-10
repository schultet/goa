package state

import (
	"fmt"
	"testing"
)

// TODO: refactor to proper test case
func TestHashIgnore(t *testing.T) {
	h := NewHash32()
	fmt.Println(h.HashIgnore([]int{1, 2, 3, 4}, 0))
	fmt.Println(h.Hash([]int{1, 2, 3, 4}))
}
