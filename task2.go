package main

import "fmt"

func matrixProd(matrixA [3][5]int, matrixB [5][4]int) [3][4]int {
	var resMatrix [3][4]int
	for i := 0; i < len(resMatrix); i++ {
		for j := 0; j < len(resMatrix[0]); j++ {
			for k := 0; k < len(matrixB); k++ {
				resMatrix[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}
	return resMatrix
}
func main() {
	matrixA := [3][5]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{4, 3, 2, 3, 4},
	}
	matrixB := [5][4]int{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{1, 3, 2, 5},
		{3, 2, 1, 4},
		{6, 4, 3, 2},
	}

	fmt.Println(matrixProd(matrixA, matrixB))
}
