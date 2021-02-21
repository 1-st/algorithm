package main

/*
{
	from:"https://leetcode-cn.com/problems/string-to-integer-atoi",
	reference:[],
	description:"atoi",
	solved:true,
}
*/

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	var max = 1<<31 - 1
	var min = -(1 << 31)
	var positive = true
	var res = 0
	var i = 0
	for str[i] == ' ' || str[i] == '0' {
		i++
		if i == len(str) {
			return 0
		}
		if str[i-1] == '0' && !(str[i] >= '0' && str[i] <= '9') {
			if i-2 >= 0 {
				if str[i-2] != '0' {
					return 0
				}
			} else {
				return 0
			}
		}
	}
	if str[i] != '+' && str[i] != '-' && !(str[i] > '0' && str[i] <= '9') {
		return 0
	}
	if str[i] == '-' {
		i++
		positive = false
	} else if str[i] == '+' {
		i++
	}
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		if res > max {
			break
		}
		res *= 10
		res += int(str[i]) - 48
		i++
	}
	if positive {
		if res <= max {
			return res
		} else {
			return max
		}
	} else {
		if -res >= min {
			return -res
		} else {
			return min
		}
	}
}
