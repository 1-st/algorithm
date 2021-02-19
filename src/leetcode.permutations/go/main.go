package _go
/*
	{
		from:"https://leetcode-cn.com/problems/permutations",
		reference:[],
		description:"全排列",
		solved:true,
	}
*/

func permute(nums []int)[][]int{
	var res [][]int
	var l = len(nums)

	var backTrack func(first int)
	backTrack = func(first int){
		if first == l	{
			res = append(res, append([]int{},nums...))
			return
		}
		for i:=first;i<l;i++{
			nums[i],nums[first] = nums[first],nums[i]
			backTrack(first+1)
			nums[i],nums[first] = nums[first],nums[i]
		}
	}

	backTrack(0)
	return res
}