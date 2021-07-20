package main

import (
	"fmt"
)

type Matrix struct {
	noOfRows,noOfColumns int
	twoD [][]int
}
func (mat Matrix) initialize()  Matrix{
	mat.twoD=make([][]int, mat.noOfRows)

	row := make([]int, mat.noOfColumns)
	for index,_ := range row{
		row[index]=0
	}

	for index, _ := range mat.twoD {
		mat.twoD[index]= make([]int, len(row))
		copy(mat.twoD[index], row)
	}
	//fmt.Println(mat)
	return mat
}

func (mat Matrix) getNumberOfRows() int {
	return mat.noOfRows
}

func (mat Matrix) getNumberOfColumns() int {
	return mat.getNumberOfColumns()
}

func (mat Matrix) setElement(i int , j int , val int)  {
	mat.twoD[i][j]=val
}

func (mat Matrix) addingMatrix(mat1 Matrix,mat2 Matrix) Matrix {
	for i:=0;i<mat1.noOfRows;i++{
		for j:=0;j<mat1.noOfColumns;j++{
			mat1.twoD[i][j]=mat1.twoD[i][j]+mat2.twoD[i][j]
		}
	}
	return mat1
}
func (mat Matrix) printMatrix()  {
	for i:=0;i<mat.noOfRows;i++{
		fmt.Println(mat.twoD[i])
	}
}

func main() {
	FirstMatrix := Matrix{noOfRows: 4,noOfColumns: 3}
	FirstMatrix=FirstMatrix.initialize()
	fmt.Println(FirstMatrix.getNumberOfRows())
	FirstMatrix.setElement(0,0,4)
	FirstMatrix.setElement(1,1,4)
	FirstMatrix.setElement(2,0,2)
	FirstMatrix.setElement(3,2,7)
	FirstMatrix.setElement(2,1,9)
	FirstMatrix.printMatrix()
}
