package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number for the multiplication table")
	fmt.Scan(&n)

	i := 1
	for i <= 10 {
		fmt.Println(n, " X ", i, " = ", n*i)
		i++

	}
}
