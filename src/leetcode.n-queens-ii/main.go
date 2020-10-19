package main

import "container/list"

/*
	{
		from:"https://leetcode-cn.com/problems/n-queens-ii",
		reference:[],
		description:"n皇后问题II",
		solved:true,
	}
*/

func totalNQueens(n int) int {
	var (
		stack     = list.New()
		ans       int
		backTrack func(s *list.List)
	)

	backTrack = func(s *list.List) {
		i := s.Len()
		if i == n {
			ans++
			return
		}
		for j := 0; j < n; j++ {
			var (
				head    = s.Front()
				p       = head
				column  = false
				leading = false
				minor   = false
			)

			var ii = 0
			for p != nil {
				jj := p.Value.(int)
				if i+j == ii+jj { //leading diagonal
					leading = true
				}
				if jj == j { //column
					column = true
				}
				if i-j == ii-jj { //minor diagonal
					minor = true
				}
				ii++
				p = p.Next()
			}

			if !column && !leading && !minor {
				s.PushBack(j)
				backTrack(s)
				s.Remove(s.Back())
			}
		}
	}
	backTrack(stack)
	return ans
}
