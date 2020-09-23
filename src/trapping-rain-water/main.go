package main

/*
	{
		from:"https://leetcode-cn.com/problems/trapping-rain-water",
		reference:[],
		description:"接雨水",
		solved:true,
	}
*/

/*
	1.由低向高扫描，填坑，如果遇到一层没有新增雨水的情况，返回
	2.scan from left to right
*/

func trap(height []int) int {
	if len(height)==0{
		return 0
	}
	var rep = make([]int, len(height))
	copy(rep, height) //save the state

	var level = 0 //current water level
	var cur = 0   //current pointer of array height
	var save = -1 // save point
	var max = 0   //the highest value in one scan

	for {
		if height[cur] >= level {
			save = cur
			level = height[cur]
			rep[cur] = height[cur]
			max = 0
		} else { // height[cur] <level
			rep[cur] = level
			if height[cur] > max {
				max = height[cur]
			}
		}
		cur++
		if cur == len(height) {
			if save+1 < len(height) {
				cur = save + 1
				level = max
			} else {
				break
			}
		}
	}
	var count = 0
	for i := 0; i < len(height); i++ {
		count += rep[i] - height[i]
	}
	return count
}
