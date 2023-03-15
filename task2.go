package main

import "fmt"

const (
	matrixACols = 5
	matrixARows = 3
	matrixBCols = 4
	matrixBRows = 5
)

func matrixProd(matrixA [matrixARows][matrixACols]int, matrixB [matrixBRows][matrixBCols]int) [matrixARows][matrixBCols]int {
	var resMatrix [matrixARows][matrixBCols]int
	for i := 0; i < matrixARows; i++ {
		for j := 0; j < matrixBCols; j++ {
			for k := 0; k < matrixBRows; k++ {
				resMatrix[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}
	return resMatrix
}
func main() {
	matrixA := [matrixARows][matrixACols]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{4, 3, 2, 3, 4},
	}
	matrixB := [matrixBRows][matrixBCols]int{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{1, 3, 2, 5},
		{3, 2, 1, 4},
		{6, 4, 3, 2},
	}

	fmt.Println(matrixProd(matrixA, matrixB))
}
