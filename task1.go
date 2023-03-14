package main

import "fmt"

func getDet(matrix [3][3]int) int {
	return matrix[0][0]*matrix[1][1]*matrix[2][2] + matrix[0][1]*matrix[1][2]*matrix[2][0] +
		matrix[0][2]*matrix[1][0]*matrix[2][1] - matrix[2][0]*matrix[1][1]*matrix[0][2] -
		matrix[0][1]*matrix[1][0]*matrix[2][2] - matrix[0][0]*matrix[1][2]*matrix[2][1]
}
func main() {

	matrix := [3][3]int{
		{1, 6, -3},
		{2, 4, 8},
		{3, -7, 0},
	}

	fmt.Println(getDet(matrix))
}
