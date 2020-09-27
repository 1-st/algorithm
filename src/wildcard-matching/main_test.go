package main

import (
	"testing"
)

type ans struct {
	s     string
	p     string
	match bool
}

func TestIsMatch(t *testing.T) {
	var cases = []ans{
		{
			s:"acdcb",
			p:"a*c?b",
			match: false,
		},
		{
			s:"c",
			p:"*?*",
			match:true,
		},
		{
			s:"abefcdgiescdfimde",
			p:"ab*cd?i*de",
			match: true,
		},
		{
			s:"aaba",
			p:"?***",
			match:true,
		},
		{
			s:"a",
			p:"a*",
			match:true,
		},
		{
			s:"aaaa",
			p:"***a",
			match:true,
		},
		{
			s:"",
			p:"*",
			match:true,
		},
		{
			s:"aa",
			p:"a*",
			match:true,
		},

		{
			s:     "aa",
			p:     "a",
			match: false,
		},
		{
			s:     "aa",
			p:     "*",
			match: true,
		},
		{
			s:     "cb",
			p:     "?a",
			match: false,
		},
		{
			s:     "adceb",
			p:     "*a*b",
			match: true,
		},
		{
			s:"",
			p:"",
			match:true,
		},

		{
			s:"a",
			p:"aa",
			match:false,
		},
	}
	for k,v:= range cases{
		if isMatch(v.s,v.p)!=v.match{
			t.Error("unmatched:",k)
		}
	}
}
