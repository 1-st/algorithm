package main

import "fmt"

type direction int

const (
	left direction = iota
	top
	right
	bottom
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var (
		leftWall   = -1
		topWall    = -1
		rightWall  = len(matrix[0])
		bottomWall = len(matrix)
		d          = top
		i          = 0
		j          = 0
		ans        []int
	)
	for bottomWall != topWall+1&&rightWall!=leftWall+1 {
		switch d {
		case top:
			for j < rightWall {
				ans = append(ans, matrix[i][j])
				j++
			}
			j--
			i++
			topWall++
			d = right
		case right:
			for i < bottomWall {
				ans = append(ans, matrix[i][j])
				i++
			}
			i--
			j--
			rightWall--
			d = bottom
		case bottom:
			for j > leftWall {
				ans = append(ans, matrix[i][j])
				j--
			}
			j++
			i--
			bottomWall--
			d = left
		case left:
			for i > topWall {
				ans = append(ans, matrix[i][j])
				i--
			}
			i++
			j++
			leftWall++
			d = top
		}
	}
	return ans
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{7},
		{9},
		{6},
	}))
}
