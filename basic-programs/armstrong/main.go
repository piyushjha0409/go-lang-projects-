package main

import (
	"fmt"
)

func main() {
	var num, temp, remainder int
	var sum int = 0
	fmt.Println("Enter the number to check :->")
	fmt.Scan(&num)

	temp = num

	//why the loop  without condition
	for {
		remainder = temp % 10
		sum += remainder * remainder * remainder
		//to loose the last digit
		temp = temp / 10

		if temp == 0 {
			break
		}
	}

	if sum == num {
		fmt.Println("The number is Armstrong!")
	} else {
		fmt.Println("The number is not Armstrong!")
	}

}
