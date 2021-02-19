package _go

import "testing"

func TestInsertInterval(t *testing.T) {
	var cases = [][2][][]int{
		{
			{
				{2,5},
				{1,3},
				{6,9},
			},
			{
				{1,5},
				{6,9},
			},
		},
		{
			{
				{6,6},
			},
			{
				{6,6},
			},
		},
		{
			{
				{6, 6},
				{3, 5},
				{12, 15},
			},
			{
				{3, 5},
				{6, 6},
				{12, 15},
			},
		},
	}
	for k, v := range cases {
		input := v[0]
		ans := v[1]
		output := insert(input[1:], input[0])
		if len(ans) != len(output) {
			t.Error("unmatched:", k)
		}
		for k, v := range output {
			if v[0] != ans[k][0] || v[1] != ans[k][1] {
				t.Error("unmatched:", k)
			}
		}
	}
}
