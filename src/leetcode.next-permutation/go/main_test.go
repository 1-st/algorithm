package _go

import "testing"

func TestNextPermutation(t *testing.T) {
	var cases = [][2][]int{
		{
			{1,5,1},{5,1,1},
		},
		{
			{}, {},
		},
		{
			{1, 3, 2}, {2, 1, 3},
		},
		{
			{4, 5, 7, 6, 4, 3, 2, 1}, {4, 6, 1, 2, 3, 4, 5, 7},
		},
		{
			{5, 4, 3, 2, 1}, {1, 2, 3, 4, 5},
		},
		{
			{1}, {1},
		},
	}

	for k, v := range cases {
		nextPermutation(v[0])
		for i := 0; i < len(v[0]); i++ {
			if v[0][i] != v[1][i] {
				t.Errorf("in case %v:%v", k, i)
			}
		}
	}
}
