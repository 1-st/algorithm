package _go

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	var cases = [][]int{
		{5, 0, 2, 2, 4, 0, 1, 0, 1, 3},
		{3, -1, 4, 2, 1, 9, 10},
		{2, 1},
		{
			3, 1, 2, 0,
		},
		{
			2, 3, 4, -1, 1,
		},
		{
			1, 7, 8, 9, 11, 12,
		},
	}
	for k, v := range cases {
		if firstMissingPositive(v[1:]) != v[0] {
			t.Error("unmatched:", k)
		}
	}
}
