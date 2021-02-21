package main

import (
	"bytes"
	"strings"
)


/*
{
	from:"https://leetcode-cn.com/problems/longest-palindromic-substring",
	reference:[],
	description:"最长回文子串",
	solved:true,
}
*/

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func longestPalindrome(s string) string {
	var strBuf bytes.Buffer
	strBuf.WriteByte('%')
	for i := 0; i < len(s); i++ {
		strBuf.WriteByte('#')
		strBuf.WriteByte(s[i])
	}
	strBuf.WriteString("#%")
	str := strBuf.String()
	var Reach = 0
	var pos = 0
	var maxIndex = 0
	var RL = make([]int, len(str))
	for i := 0; i < len(str); i++ {
		if i < Reach {
			RL[i] = min(RL[2*pos-i], Reach-i)
		} else {
			RL[i] = 1
		}
		for i-RL[i] >= 0 && i+RL[i] < len(str) && str[i-RL[i]] == str[i+RL[i]] {
			RL[i]++
		}
		if RL[i]+i-1 > Reach {
			Reach = RL[i] + i - 1
			pos = i
		}
		if RL[maxIndex] < RL[i] {
			maxIndex = i
		}
	}
	var res = ""
	for i := 0; i < RL[maxIndex]; i++ {
		if i == 0 {
			res = res + string(str[maxIndex])
		} else {
			res = string(str[maxIndex-i]) + res + string(str[maxIndex+i])
		}
	}
	return strings.Replace(
		strings.Replace(
			res, "#", "", -1),
		"%", "", -1)
}
