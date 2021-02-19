package main

import "fmt"

/*
{
	from:"https://leetcode-cn.com/problems/jump-game",
	reference:[],
	description:"跳跃游戏",
	solved:true,
}
*/

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	var (
		reach = 0
		cur   = 0
	)
	for cur < len(nums)-1 && cur <= reach {
		if nums[cur]+cur > reach {
			reach = nums[cur] + cur
		}
		cur++
		if reach >= len(nums)-1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(canJump([]int{0, 2, 3}))
}
