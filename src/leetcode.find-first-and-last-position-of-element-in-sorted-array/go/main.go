package _go

/*
{
	from:"https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/",
	reference:[
		"https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/solution/"
	],
	description:"寻找非严格单调数组中的target区间",
	solved:true,
}
*/

/*
	1.对两侧进行单侧区间的二分查找

	tip:注意第一次单侧区间查找返回的结果可能出错,并且需要检查出错下表越界的可能
*/

func halfBinarySearch(nums []int, target int, left bool) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target { //由于要寻找区间,所以不结束
			if left { //向左侧查找
				j = mid - 1
			} else {
				i = mid + 1
			}
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	if left {
		return i
	} else {
		return j
	}
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	lIndex := halfBinarySearch(nums, target, true)
	if lIndex == len(nums) {
		return res
	}
	if nums[lIndex] != target {
		return res
	}
	rIndex := halfBinarySearch(nums, target, false)
	res[0], res[1] = lIndex, rIndex
	return res
}
