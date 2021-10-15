package main

import "fmt"

func main() {
	var matrix [3][3]int = [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(matrix)
	matrix[0][0] = 9 // replacing particular item in array
	fmt.Println(matrix)
}
