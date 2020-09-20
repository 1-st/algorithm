package main

import "container/list"

/*
	{
		from:"https://leetcode-cn.com/problems/longest-valid-parentheses",
		reference:[],
		description:"最长有效括号长度",
		solved:true,
	}
*/

/*
	1.额外维护一个相同长度的bool数组,保存相应位置括号是否已经被匹配,初始全部为false,
	从左到右将括号带着位置一起入栈，一旦完成匹配，出栈并且将相应位置的bool数组设为true
	最后扫描一遍bool数组,获取最长有效长度,时间复杂度O(n)

	tips:使用container/list 表示栈
*/

type Item struct {
	Pos  int
	Left bool //(
}

func longestValidParentheses(s string) int {
	len := len(s)
	if len == 0 || len == 1 {
		return 0
	}

	valid := make([]bool, len) //false

	var stack = list.New()

	for i := 0; i < len; i++ {
		if stack.Len() == 0 {
			var left = true
			if s[i] == ')' {
				left = false
			}
			stack.PushBack(Item{
				Pos:  i,
				Left: left,
			})
		} else {
			var last = stack.Back().Value.(Item)
			if s[i] == ')' && last.Left {
				valid[i] = true
				valid[last.Pos] = true
				stack.Remove(stack.Back())
			} else {
				var left = true
				if s[i] == ')' {
					left = false
				}
				stack.PushBack(Item{
					Pos:  i,
					Left: left,
				})
			}
		}
	}

	var max = 0
	var cur = 0
	for i := 0; i < len; i++ {
		if valid[i] {
			cur++
			if cur > max {
				max = cur
			}
		} else {
			cur = 0
		}
	}

	return max
}
