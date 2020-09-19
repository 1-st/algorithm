package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	var cases = [][]int{
		//target result array
		{0, 4, /**/ 4, 5, 6, 7, 0, 1, 2},
		{3, -1, /**/ 4, 5, 6, 7, 0, 1, 2},
		{2, -1},
		{3, -1, /**/ 2},
		{0, 0, /**/ 0},
	}
	for k, v := range cases {
		if search(v[2:], v[0]) != v[1] {
			t.Errorf("unmatched %v", k)
			return
		}
	}
}
