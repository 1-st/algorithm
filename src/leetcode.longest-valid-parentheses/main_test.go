package main

import "testing"

func TestLongestValidParentheses(t *testing.T) {

	cases := map[string]int{
		"()":                    2,
		"()((((())))()((()())(": 10,
		"":                      0,
		"))))))))))":            0,
	}
	for k, v := range cases {
		if longestValidParentheses(k) != v {
			t.Errorf("unmatched %v", k)
			return
		}
	}
}
