package _go

/*
{
	from:"https://leetcode-cn.com/problems/jump-game-ii",
	reference:["https://leetcode-cn.com/problems/jump-game-ii/solution"],
	description:"跳跃游戏 II",
	solved:true,
}
*/

/*
	tips:使用贪心算法比较好
*/

func jump(nums []int) int {
	var cost = make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		var min = cost[i+1] + 1
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			if cost[i+j]+1 < min {
				min = cost[i+j] + 1
			}
		}
		cost[i] = min
	}
	return cost[0]
}
