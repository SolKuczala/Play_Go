package main

import (
	"fmt"
	"strconv"
)

func main() {
	var matrix [3][3]string //[[inl, , ],[ , , ],[ , , ]]
	fmt.Printf("%+v\n", matrix)

	for i := 0; i < len(matrix); i++ {
		//me trae cada aarray de matrix
		for j := 0; j < len(matrix[i]); j++ {
			//me trae 1 valor de cada array
			matrix[i][j] = strconv.Itoa(i) + strconv.Itoa(j)
		}
	}
	fmt.Printf("%+v", matrix)
	//for i, array := range matrix {
	//	//cada array de matrix
	//	for i := 0; i < count; i++ {
	//
	//	}

}
