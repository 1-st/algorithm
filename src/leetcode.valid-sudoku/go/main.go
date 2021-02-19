package _go

/*
{
	from:"https://leetcode-cn.com/problems/valid-sudoku",
	reference:[],
	description:"有效的数独",
	solved:true,
}
*/

/*
	1.一遍遍历，同时存储
*/

func isValidSudoku(board [][]byte) bool {
	var row [9][9]bool
	var column [9][9]bool
	var box [9][9]bool
	l := len(board)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if board[i][j]=='.'{
				continue
			}
			var n = board[i][j]-'1'
			if !row[i][n]{
				row[i][n] = true
			}else{
				return false
			}
			if !column[j][n]{
				column[j][n] = true
			}else{
				return false
			}
			if !box[i/3*3+j/3][n]{
				box[i/3*3+j/3][n] = true
			}else{
				return false
			}
		}
	}
	return true
}
