package main

import "fmt"

/*
	{
		from:"https://leetcode-cn.com/problems/unique-paths",
		reference:[],
		description:"不同路径",
		solved:true,
	}
*/

func uniquePaths(m int, n int) int {
	var C = func(a, b int) int {
		if a < b {
			a, b = b, a
		}
		if b == 0 || a == b {
			return 1
		}
		if b > (a - b) {
			b = a - b
		}
		var sum = 1
		for i := 1; i <= b; i++ {
			sum *= a - (i - 1)
			sum /= i
		}
		return sum
	}
	return C(n-1, m+n-2)
}

//
//func uniquePaths(m int, n int) int {
//	var grid = make([][]bool, m)
//	for i := 0; i < m; i++ {
//		grid[i] = make([]bool, n)
//	}
//	var back func(i, j int) int
//	back = func(i, j int) int {
//		if i == m-1 && j == n-1 {
//			return 1
//		}
//		sum := 0
//		if i+1 < m && !grid[i+1][j] {
//			grid[i+1][j] = true
//			sum += back(i+1, j)
//			grid[i+1][j] = false
//		}
//
//		if j+1 < n && !grid[i][j+1] {
//			grid[i][j+1] = true
//			sum += back(i, j+1)
//			grid[i][j+1] = false
//		}
//
//		return sum
//	}
//	return back(0,0)
//}

func main() {
	fmt.Println(uniquePaths(7, 3))
}
