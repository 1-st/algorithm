package _go

import "testing"

func TestMultiply(t *testing.T) {
	var cases = [][3]string{
		{"123", "456", "56088"},
		{"2", "3", "6"},
		{"0", "0", "0"},
	}
	for k, v := range cases {
		if multiply(v[0], v[1]) != v[2] {
			t.Error("unmatched:", k)
		}
	}
}
