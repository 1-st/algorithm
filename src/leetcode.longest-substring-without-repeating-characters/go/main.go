package main

/*
	{
		from:"https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/",
		reference:[],
		description:"无重复字符的最长子串",
		solved:true,
	}
*/


func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen==0{
		return 0
	}
	i,j := 0,0
	maxLen := 1
	dict := make(map[byte]int)
	dict[s[i]] = i
	for sLen-i>maxLen&&j<sLen-1{
		j++
		if iPre,ok:=dict[s[j]];!ok||(ok&&iPre<i) {
			dict[s[j]]=j
			if maxLen<j-i+1{
				maxLen = j-i+1
			}
		}else{
			i = dict[s[j]]+1
			dict[s[j]] = j
		}
	}
	return maxLen
}

