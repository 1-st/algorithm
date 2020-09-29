package main

import (
	"fmt"
	"sort"
)

type Interval [][]int

func (p Interval) Len() int           { return len(p) }
func (p Interval) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p Interval) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func merge(intervals [][]int) [][]int {
	sort.Sort(Interval(intervals))
	cur := 0
	for cur < len(intervals)-1 {
		for intervals[cur][1] >= intervals[cur+1][0] { //merge
			if intervals[cur][1]<intervals[cur+1][1]{
				intervals[cur+1][0] = intervals[cur][0]
				intervals = append(intervals[:cur], intervals[cur+1:]...)
			}else{
				intervals = append(intervals[:cur+1], intervals[cur+2:]...)
			}
			if len(intervals) == cur+1 {
				return intervals
			}
		}
		cur++
	}
	return intervals
}

func main() {
	fmt.Println(merge([][]int{
		{
			2,3,
		},
		{
			2, 2,
		},
		{
			3,3,
		},
		{
			1,3,
		},
		{
			5,7,
		},
		{
			2,2,
		},
		{
			4,6,
		},
	}))
}
