package main

import (
	"container/list"
	"fmt"
)

func solveNQueens(n int) [][]string {
	var (
		stack     = list.New()
		ans       [][]string
		backTrack func(s *list.List)
	)

	backTrack = func(s *list.List) {
		i := s.Len()
		if i == n {
			ans = append(ans, append([]string{}, export(s)...))
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
				if jj == j {//column
					column = true
				}
				if i-j == ii-jj{//minor diagonal
					minor = true
				}
				ii++
				p = p.Next()
			}

			if !column&&!leading&&!minor {
				s.PushBack(j)
				backTrack(s)
				s.Remove(s.Back())
			}
		}
	}
	backTrack(stack)
	return ans
}

func export(queen *list.List) []string {
	var (
		ans []string
		l   = queen.Len()
		p   = queen.Front()
	)

	for p != nil {
		var (
			v   = p.Value.(int)
			str = ""
		)
		for i := 0; i < v; i++ {
			str += "."
		}
		str += "Q"
		for i := 0; i < l-v-1; i++ {
			str += "."
		}
		ans = append(ans, str)
		p = p.Next()
	}

	return ans
}

func main() {

	fmt.Println(len(solveNQueens(8)))
}
