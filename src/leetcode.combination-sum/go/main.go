package _go

import "sort"

/*
{
	from:"https://leetcode-cn.com/problems/combination-sum",
	reference:[
		"https://leetcode-cn.com/problems/combination-sum/solution/"
	],
	description:"组合总和",
	solved:true,
}
*/

/*
	1.深度为len(candidates),candidates数组每个项有两种选择，选择或者跳过
	选择之前先排序，方便剪枝

	tips:注意append slice的时候需要拆分slice为单个元素再重新组合,不然只是添加了引用
*/

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var comb []int
	var res [][]int
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if index == len(candidates) {
			return
		} else if target == 0 {
			res = append(res, append([]int{}, comb...)) //*
			return
		}else if target <candidates[index]{ //当前项不满足，未来更大的项也不会满足
			return
		}
		//跳过
		dfs(target, index+1)
		if target-candidates[index] >= 0 {
			comb = append(comb, candidates[index])
			dfs(target-candidates[index], index)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return res
}
