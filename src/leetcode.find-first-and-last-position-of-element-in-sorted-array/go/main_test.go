package _go

import "testing"

func TestSearchRange(t *testing.T) {
	var cases = [][]int{
		// result target array
		{-1,-1,3,2,2},
		{3, 4, 8, 5, 7, 7, 8, 8, 10},
		{-1, -1, 6, 5, 7, 7, 8, 8, 10},
		{0,0,5,5},
		{-1,-1,5},
	}
	for k, v := range cases {
		res := searchRange(v[3:], v[2])
		if res[0] != v[0] || res[1] != v[1] {
			t.Errorf("unmatched %v", k)
			return
		}
	}
}
