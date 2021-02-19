package main

import (
	"fmt"
)

/*
{
	from:"https://leetcode-cn.com/problems/minimum-path-sum",
	reference:[],
	description:"Minimum Path Sum",
	solved:true,
	language:"Go",
}
*/
func min(x ,y int)int{
	if x<y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	var dp = make([][]int,len(grid))
	for i:=0;i<len(grid);i++{
		dp[i] = make([]int,len(grid[0]))
	}
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if i==0&&j==0{
				dp[i][j] = grid[i][j]
			}else if i==0&&j>0{
				dp[i][j] = grid[i][j]+dp[i][j-1]
			}else if j==0&&i>0{
				dp[i][j] = grid[i][j]+dp[i-1][j]
			}else if i>0&&j>0{
				dp[i][j] = grid[i][j] + min(dp[i-1][j],dp[i][j-1])
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func main() {
	grid1 := [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	fmt.Println(minPathSum(grid1))
}
