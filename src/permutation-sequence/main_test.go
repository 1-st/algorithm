package main

import "testing"

type Result struct {
	N   int
	K   int
	Str string
}

func TestGetPermutation(t *testing.T) {
	var cases = []Result{
		{
			N:1,
			K:1,
			Str:"1",
		},
		{
			N:5,
			K:1,
			Str:"12345",
		},
		{
			N:5,
			K:2,
			Str:"12354",
		},
		{
			N:   4,
			K:   10,
			Str: "2341",
		},
		{
			N:   4,
			K:   9,
			Str: "2314",
		},
		{
			N:   3,
			K:   3,
			Str: "213",
		},
	}
	for k, v := range cases {
		if getPermutation(v.N, v.K) != v.Str {
			t.Error("unmatched:", k)
		}
	}
}
