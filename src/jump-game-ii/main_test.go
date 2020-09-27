package main

import "testing"

func TestJump(t *testing.T) {
	var cases = [][]int{
		{
			2, 2, 3, 1, 1, 4,
		},
		{
			3, 4, 5, 1, 1, 2, 8, 9, 2,
		},
	}
	for k, v := range cases {
		if jump(v[1:]) != v[0] {
			t.Error("unmatched:", k)
		}
	}
}
