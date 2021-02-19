package _go

import "sort"

/*
{
	from:"https://leetcode-cn.com/problems/combination-sum-ii",
	reference:[
		"https://leetcode-cn.com/problems/combination-sum-ii/solution/"
	],
	description:"组合总和II",
	solved:true,
}
*/

/*
	1.由于元素不可重复选取,所以相较上题可以以另一种方式递归,

	tips:注意第二处减枝不要习惯性加上return
 */


func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var path []int
	var ans [][]int
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
		}
		for i := idx; i < len(candidates); i++ {
			if target < candidates[i] {
				return
			}
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(target-candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(target, 0)
	return ans
}

//func combinationSum2(candidates []int, target int) [][]int {
//	sort.Ints(candidates)
//	var comb []int
//	var res [][]int
//	var dfs func(target, index int)
//	dfs = func(target, index int) {
//		if index == len(candidates) {
//			return
//		} else if target == 0 {
//			res = append(res, append([]int{}, comb...)) //*
//			return
//		} else if target < candidates[index] { //当前项不满足，未来更大的项也不会满足
//			return
//		}else if index>=0&&candidates[index]==candidates[index-1]{
//			return
//		}
//		//跳过
//		dfs(target, index+1)
//		if target-candidates[index] >= 0 {
//			comb = append(comb, candidates[index])
//			dfs(target-candidates[index], index+1) //虽然防止了元素的重复,但是没有做到结果去重
//			comb = comb[:len(comb)-1]
//		}
//	}
//	dfs(target, 0)
//	return res
