package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{50, 60, 70, 80, 90}
	fmt.Println(num)

	//check array of the integer is sorted
	if sort.IntsAreSorted(num) == false {
		sort.Ints(num)
	}

	text := []string{"Uk", "Usa", "India", "Germany"}
	fmt.Println(text)

	if sort.StringsAreSorted(text) == false {
		sort.Strings(text)
	}
	fmt.Println(text)
	fmt.Println(sort.SearchStrings(text, "India"))
}
