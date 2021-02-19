package _go

import "testing"

func TestTrap(t *testing.T) {
	var cases = [][]int{
		{
			6, 0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1,
		},
		{
			0,0,
		},
		{
			8,8,0,8,
		},
		{
			9,8,0,8,0,1,
		},
		{0},
	}
	for k, v := range cases {
		if trap(v[1:]) != v[0] {
			t.Error("unmatched:",k)
		}
	}
}
