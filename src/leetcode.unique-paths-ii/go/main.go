package main

import "fmt"

/*
{
	from:"https://leetcode-cn.com/problems/lru-cache",
	reference:[],
	description:"LRU cache",
	solved:true,
	language:"Go",
}
*/

/*
	使用动态规划避免深度优先算法
	TODO 使用原二维数组存储动态规划的值
*/

//func dfs(grid [][]int, x, y int) int {
//	if x >= len(grid) || y >= len(grid[0]) {
//		return 0
//	}
//	if grid[x][y]==1{
//		return 0
//	}
//	if x == len(grid)-1 && y == len(grid[0])-1 {
//		return 1
//	}
//	return dfs(grid, x+1, y) + dfs(grid, x, y+1)
//}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	var grid = make([][]int, len(obstacleGrid))
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, len(obstacleGrid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if obstacleGrid[i][j] == 0 {
				if i == 0 && j != 0 {
					if obstacleGrid[i][j-1] == 0 {
						grid[i][j] = grid[i][j-1]
					}
				} else if j == 0 && i != 0 {
					if obstacleGrid[i-1][j] == 0 {
						grid[i][j] = grid[i-1][j]
					}
				} else if i == 0 && j == 0 {
					grid[i][j] = 1
				} else {
					if obstacleGrid[i-1][j] == 0 {
						grid[i][j] += grid[i-1][j]
					}
					if obstacleGrid[i][j-1] == 0 {
						grid[i][j] += grid[i][j-1]
					}
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	grid1 := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(grid1))

	grid2 := [][]int{
		{0, 1},
		{0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(grid2))
}
