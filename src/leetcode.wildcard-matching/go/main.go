package _go

/*
{
	from:"https://leetcode-cn.com/problems/wildcard-matching",
	reference:["https://leetcode-cn.com/problems/wildcard-matching/solution"],
	description:"通配符匹配",
	solved:false,
}
*/

/*
	1.逐个扫描 false
	2.动态规划
*/

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true               //edge case 1
	for i := 1; i <= len(p); i++ { //edge case 2
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}

//func compare(s string, p string) bool {
//	var i, j = 0, 0
//	for i < len(s) && j < len(p) {
//		if s[i] == p[j] || p[j] == '?' {
//			i++
//			j++
//		} else {
//			return false
//		}
//	}
//	if i == len(s) && j == len(p) {
//		return true
//	}
//	return false
//}
//
//func isMatch(s string, p string) bool {
//	if len(p) == 0 {
//		if len(s) == 0 {
//			return true
//		}
//		return false
//	}
//
//	var pp, ps = 0, 0
//
//	for pp != len(p) && ps != len(s) {
//		switch p[pp] {
//		case '*':
//			if pp+1 == len(p) {
//				return true
//			}
//			//get stop token
//			var left = pp + 1
//			for left < len(p) {
//				if p[left] != '*' {
//					break
//				}
//				left++
//			}
//			if left == len(p) {
//				return true
//			}
//			var right = left
//			for right < len(p) {
//				if p[right] == '*' {
//					break
//				}
//				right++
//			}
//			for i := ps; i < len(s); i++ {
//				if i+right-left <= len(s) && compare(s[i:i+right-left], p[left:right]) {
//					if isMatch(s[i:], p[left:]) {
//						return true
//					}
//				}
//			}
//			return false
//		case '?':
//			pp++
//			ps++
//		default: //a-z
//			if p[pp] != s[ps] {
//				return false
//			} else {
//				pp++
//				ps++
//			}
//		}
//	}
//
//	if pp == len(p) && ps == len(s) {
//		return true
//	}
//	if ps == len(s) {
//		//p end with several *
//		for i := pp; i < len(p); i++ {
//			if p[i] != '*' {
//				return false
//			}
//		}
//		return true
//	}
//	return false
//}
