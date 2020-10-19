package main

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append([][]int{}, newInterval)
	}
	var (
		start = newInterval[0]
		end   = newInterval[1]
		left  = -1
		right = len(intervals)
		lIn   = false
		rIn   = false
	)
	var i = 0
	for {
		if lIn && rIn {
			break
		}
		if intervals[i][0] <= start && !lIn {
			left = i
			lIn = intervals[i][1] >= start
		}
		if intervals[len(intervals)-1-i][1] >= end && !rIn {
			right = len(intervals) - 1 - i
			rIn = intervals[len(intervals)-1-i][0] <= end
		}
		if i == len(intervals)-1 {
			break
		}
		i++
	}
	//right
	if left == len(intervals)-1 && !lIn {
		return append(intervals, append([]int{}, newInterval...))
	}
	//left
	if right == 0 && !rIn {
		return append(append([][]int{}, newInterval), intervals...)
	}
	//inner
	if left+1==right&&!lIn&&!rIn{
		intervals = append(intervals,[]int{})
		copy(intervals[left+2:],intervals[left+1:])
		intervals[left+1] = newInterval
		return intervals
	}
	if !lIn {
		left++
	}
	if !rIn {
		right--
	}
	var mixed = make([]int, 2)
	if start < intervals[left][0] {
		mixed[0] = start
	} else {
		mixed[0] = intervals[left][0]
	}
	if end > intervals[right][1] {
		mixed[1] = end
	} else {
		mixed[1] = intervals[right][1]
	}

	l:= len(intervals)-(right-left)
	intervals = append(intervals,[]int{})
	copy(intervals[left+1:],intervals[right+1:])
	intervals[left] = mixed
	return intervals[:l]
}

