package _go

/*
{
	from:"https://leetcode-cn.com/problems/sudoku-solver",
	reference:[
		"https://leetcode-cn.com/problems/sudoku-solver/solution/"
	],
	description:"解数独",
	solved:true,
}
*/

/*
	尝试用位运算优化空间效率
*/

type Solution struct {
	board *[][]byte
	Empty []int  //                     000000000(i)         000000000(j)
	State [9]int // 00000 000000000(row) 000000000(column) 000000000(box)
	Stop  bool
}

func (s *Solution) dfs(solved int) {
	if solved == len(s.Empty) {
		s.Stop = true
		return
	}
	i, j := s.Empty[solved]>>9, s.Empty[solved]&0x1ff
	for k := 0; k < 9 && !s.Stop; k++ {
		if s.canSet(i, j, k) {
			s.setState(i, j, k, true)
			(*s.board)[i][j] = byte(k + '1')
			s.dfs(solved + 1)
			s.setState(i, j, k, false)
		}
	}
}

func (s *Solution) canSet(i, j, n int) bool {
	if s.State[i]&(1<<n) == 0 && s.State[j]&(1<<(n+9)) == 0 && s.State[(i/3)*3+j/3]&(1<<(n+18)) == 0 {
		return true
	}
	return false
}

func (s *Solution) setState(i, j, n int, ok bool) {
	if ok {
		s.State[i] |= 1 << n
		s.State[j] |= 1 << (n + 9)
		s.State[(i/3)*3+j/3] |= 1 << (n + 18)
	} else {
		s.State[i] &^= 1 << n
		s.State[j] &^= 1 << (n + 9)
		s.State[(i/3)*3+j/3] &^= 1 << (n + 18)
	}
}

func solveSudoku(board [][]byte) {
	var s = Solution{
		board: &board,
		Empty: make([]int, 0),
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == '.' {
				var m = 0
				m |= j
				m |= i<<9
				s.Empty = append(s.Empty, m)
			} else {
				n := int(board[i][j] - '1')
				s.setState(i, j, n, true)
			}
		}
	}
	s.dfs(0)
}
