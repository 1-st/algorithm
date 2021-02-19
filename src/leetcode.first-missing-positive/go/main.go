package _go

/*
	{
		from:"https://leetcode-cn.com/problems/first-missing-positive",
		reference:[],
		description:"缺失的第一个正数",
		solved:true,
	}
*/

/*
	1.先快速排序数组，然后从最小数字开始（如果有的话），中断直接输出结果，
	如果最小数字大于1，直接输出1 false 时间复杂度太大

	2.复用数组存储遍历过的数据，将数存储到下标对应的数字单位上，与原来的数字交换，如没有可以交换的位置，将数字置为0
	遍历

	tips:如果遇到未就位的数据，不能简单的交换，需要循环到遇到已就位的数据
*/

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > len(nums) { //out of range
			nums[i] = 0
		} else {
			j := nums[i] - 1
			for j>=0&&nums[j] != j+1 {
				temp := nums[j] - 1
				nums[j] = j + 1
				j = temp
			}
		}
	}
	var exp = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] > exp {
			return exp
		} else if nums[i] == exp {
			exp++
		}
	}
	return exp
}

//func firstMissingPositive(nums []int) int {
//	if len(nums) == 0 {
//		return 1
//	}
//	sort.Ints(nums)
//	if nums[0] > 1 {
//		return 1
//	}
//	var last = nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i] > 1 {
//			if nums[i]-last > 1 {
//				return nums[i] - 1
//			}
//		}
//		last = nums[i]
//	}
//	return nums[len(nums)-1] + 1
//}
