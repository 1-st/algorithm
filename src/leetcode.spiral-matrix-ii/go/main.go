package main

import "fmt"

type direction int

const (
	left direction = iota
	top
	right
	bottom
)

func generateMatrix(n int) [][]int {

	var matrix = make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}
	var (
		leftWall   = -1
		topWall    = -1
		rightWall  = n
		bottomWall = n
		d          = top
		i          = 0
		j          = 0
		x = 1
	)
	for bottomWall != topWall+1 && rightWall != leftWall+1 {
		switch d {
		case top:
			for j < rightWall {
				matrix[i][j] = x
				x++
				j++
			}
			j--
			i++
			topWall++
			d = right
		case right:
			for i < bottomWall {
				matrix[i][j] = x
				x++
				i++
			}
			i--
			j--
			rightWall--
			d = bottom
		case bottom:
			for j > leftWall {
				matrix[i][j] = x
				x++
				j--
			}
			j++
			i--
			bottomWall--
			d = left
		case left:
			for i > topWall {
				matrix[i][j] = x
				x++
				i--
			}
			i++
			j++
			leftWall++
			d = top
		}
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
}
