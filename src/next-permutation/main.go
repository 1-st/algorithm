package main

/*
	{
		from:"https://leetcode-cn.com/problems/next-permutation",
		reference:[],
		description:"实现获取下一个排列的函数",
		solved:true,
	}
*/

/*
	1.自右向左冒泡排序,如果len-1个数字均无法向左移动,则使用双指针倒转 false
	2.从右向左维护一个非严格单调递减序列，遇到第一个不满足情况的数i，在序列中自左向右查找大于i的数中最小那个,
	与数i替换,在这之后将新数i之后的数字使用双指针法倒序

	tips:
		小心将O(n) 代码写出嵌套for 循环
*/

func nextPermutation(nums []int) {
	len := len(nums)

	if len == 0 || len == 1 {
		return
	}

	i := len - 2
	for ; i >= 0; i--{
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i>=0 {
		for j := i + 1; j < len; j++ {
			if nums[j] <= nums[i] {
				nums[i], nums[j-1] = nums[j-1], nums[i]
				break
			} else if j == len-1 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}

	for k := i + 1; k <= (i+len)/2; k++ {
		nums[k], nums[len-(k-i)] = nums[len-(k-i)], nums[k]
	}
}
