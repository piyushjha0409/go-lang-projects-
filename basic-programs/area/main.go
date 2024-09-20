package main

import "fmt"

var area int

func main() {
	var l, b int
	fmt.Println("Enter the length and breadth of the rctangle: ")
	fmt.Scan(&l, &b)
	area = l * b
	fmt.Println("The area of the rectangle is: ", area)
}
