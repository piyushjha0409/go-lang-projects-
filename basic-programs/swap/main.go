package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(strings.Compare("India", "Indiana"))
	// fmt.Println(strings.Compare("Indiana", "India"))
	// fmt.Println(strings.Compare("India", "India"))
	fmt.Println(strings.Compare("Indiana", "India"))

	var i, sum int
	for i = 1; i < 100; i++ {
		sum += i
	}

	fmt.Println("Here is the answer %d\n", sum)
}
