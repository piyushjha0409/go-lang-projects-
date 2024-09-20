package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the no of lines for the pascals triangle: ")
	fmt.Scan(&num)
	var curr int = 1

	for i := 0; i < num; i++ {

		//print the white spaces
		for j := 1; j <= num-i; j++ {
			fmt.Print(" ")
		}

		//now print the sum in the middle
		for k := 0; k <= i; k++ {

			if k == 0 || i == 0 {
				curr = 1
			} else {
				curr = curr * (i - k + 1) / k
			}

			fmt.Printf(" %d", curr)
		}
		fmt.Println("")
	}
}
