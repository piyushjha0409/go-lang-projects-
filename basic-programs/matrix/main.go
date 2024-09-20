package main

import "fmt"

func main() {
	var matrix1 [3][3]int
	var matrix2 [3][3]int
	var sum [3][3]int
	var row, col int

	fmt.Println("Print the number of rows: ")
	fmt.Scan(&row)
	fmt.Println("Print the number of columns: ")
	fmt.Scan(&col)

	fmt.Println()

	fmt.Println(" =========== Matrix 1 ============")
	fmt.Println()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter the elements for the  matrix 1  %d %d  ->", i+1, j+1)
			fmt.Scanln(&matrix1[i][j])
		}
	}

	fmt.Println(" =========== Matrix 1 ============")
	fmt.Println()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter the elements for the  matrix 2	%d %d  ->", i+1, j+1)
			fmt.Scanln(&matrix2[i][j])
		}
	}

	//for sum
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}

	//for printing the sum matrix

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ", sum[i][j])
			if j == col-1 {
				fmt.Println("")
			}
		}
	}

}
