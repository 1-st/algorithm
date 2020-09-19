package main

/*!
{
	from:"https://leetcode-cn.com/problems/search-in-rotated-sorted-array",
	reference:[
		"https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/"
	],
	description:"搜索旋转排列数组",
	solved:true,
}
*/

/*
	1.由于任意两两划分数组,总会出现一段有序另一段左端大右端小,所以可以每次二分划分数组,
	然后判断target是否在有序段,如果不在有序段,则继续划分,如果在有序段,直接在有序段二分查找

	tips:仔细观察数组性质
		写复杂了!!!!!!,代码可以在一个for循环中完成
*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	var i, j = 0, len(nums) - 1
	var mid = (i + j) / 2

	if nums[i] <= nums[j] { //全有序
		return binarySearch(target, i, j, &nums)
	}

	if nums[i] <= nums[mid] { //左端有序
		if target >= nums[i] && target <= nums[mid] { //target在左端
			j = mid
			return binarySearch(target, i, j, &nums)
		} else {
			var index = search(nums[mid+1:], target)
			if index == -1 {
				return -1
			} else {
				return index + mid + 1
			}
		}
	} else {                                          //右端有序
		if target >= nums[mid] && target <= nums[j] { //target在右端
			i = mid
			return binarySearch(target, i, j, &nums)
		} else {
			return search(nums[:mid], target)
		}
	}
}

func binarySearch(target int, i int, j int, nums *[]int) int {
	for i <= j {
		mid := (i + j) / 2
		if target == (*nums)[mid] {
			return mid
		} else if target < (*nums)[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}

//func search_better(nums []int, target int) int {
//	low,high:=0,len(nums)-1
//	for low < high - 1{
//		mid:=(high+low)/2
//		a,b,c:=nums[low],nums[mid],nums[high]
//		if b == target{
//			return mid
//		}
//		switch {
//		// 上面两个case在递增序列找同二分查找
//		case target>=a && target<=b:
//			high = mid
//		case target>=b && target<=c:
//			low = mid
//		// 如果在递增的区间没有找到
//		// 那么只可能在递减的区间中
//		// 由于是旋转数组，只可能有一个递减区间
//		case a>b:
//			high = mid
//		case b>c:
//			low  = mid
//		default:
//			return -1
//		}
//	}
//	if nums[low] == target{
//		return low
//	}
//	if nums[high] == target{
//		return high
//	}
//	return -1
//}
//
