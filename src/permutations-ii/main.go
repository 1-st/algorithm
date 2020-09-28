package main

import (
	"fmt"
	"sort"
)

/*
	{
		from:"https://leetcode-cn.com/problems/permutations-ii",
		reference:[],
		description:"全排列 II",
		solved:true,
	}
*/

type Set struct {
	content map[interface{}]bool
}

func NewSet() *Set {
	return &Set{
		content: map[interface{}]bool{},
	}
}
func (s *Set) Add(value interface{}) {
	s.content[value] = true
}

func (s *Set) Has(value interface{}) bool {
	_, ok := s.content[value]
	return ok
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var l = len(nums)
	sort.Ints(nums)
	var backTrack func(first int)
	backTrack = func(first int) {
		if first == l {
			res = append(res, append([]int{}, nums...))
			return
		}
		var set = NewSet()
		set.Add(nums[first])
		for i := first; i < l; i++ {
			if i != first {
				if set.Has(nums[i]) {
					continue
				}else{
					set.Add(nums[i])
				}
			}
			nums[i], nums[first] = nums[first], nums[i]
			backTrack(first + 1)
			nums[i], nums[first] = nums[first], nums[i]
		}
	}

	backTrack(0)
	fmt.Println(res)
	return res
}

func main() {
	permuteUnique([]int{1,1,2,2})
}
