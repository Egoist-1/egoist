package tools

import (
	"fmt"
	"testing"
)

type SrcSlice struct {
	n   string
	age int
	no  string
}
type DstSlice struct {
	name string
	i    int
}

func TestSliceMap(t *testing.T) {
	src := []SrcSlice{{
		n:   "1",
		age: 0,
		no:  "1",
	}, {
		n:   "2",
		age: 2,
		no:  "2",
	}, {
		n:   "3",
		age: 3,
		no:  "3",
	}}
	slices := Map[SrcSlice, DstSlice](src, func(idx int, s SrcSlice) DstSlice {
		return DstSlice{
			name: s.n,
			i:    s.age,
		}
	})
	fmt.Println(slices)
}
