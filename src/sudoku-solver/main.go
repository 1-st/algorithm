package main

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

type Solution struct {
	board  *[][]byte
	Empty  [][2]int
	Row    [9][9]bool
	Column [9][9]bool
	Box    [9][9]bool
	Stop bool
}

func (s *Solution) bfs(solved int) {
	if solved == len(s.Empty) {
		s.Stop = true
		return
	}
	i, j := s.Empty[solved][0], s.Empty[solved][1]
	for k := 0; k < 9&&!s.Stop; k++ {
		if s.canSet(i, j, k) {
			s.setState(i, j, k, true)
			(*s.board)[i][j] = byte(k + '1')
			s.bfs(solved + 1)
			s.setState(i, j, k, false)
		}
	}
}

func (s *Solution) canSet(i, j, n int) bool {
	if !s.Row[i][n] && !s.Column[j][n] && !s.Box[(i/3)*3+j/3][n] {
		return true
	}
	return false
}

func (s *Solution) setState(i, j, n int, ok bool) {
	s.Row[i][n] = ok
	s.Column[j][n] = ok
	s.Box[(i/3)*3+j/3][n] = ok
}

func solveSudoku(board [][]byte) {
	var s = Solution{
		board: &board,
		Empty: make([][2]int, 0),
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == '.' {
				s.Empty = append(s.Empty, [2]int{i, j})
			} else {
				n := int(board[i][j] - '1')
				s.setState(i, j, n, true)
			}
		}
	}
	s.bfs(0)
}
