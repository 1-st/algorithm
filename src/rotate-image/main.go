package main

import "fmt"

/*
	{
		from:"https://leetcode-cn.com/problems/rotate-image",
		reference:[],
		description:"旋转图像",
		solved:true,
	}
*/

/*
	1.对角旋转，每行倒序
*/

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}
	}
}

func main() {
	matrix:= [][]int{
		{
			1, 2, 3,
		},
		{
			4, 5, 6,
		},
		{
			7, 8, 9,
		},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
