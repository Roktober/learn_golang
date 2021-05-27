package main

import (
	"fmt"
	"premetiv_sort/sort"
)

func main() {
	inputArr := [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}
	arrSlice := make([]int, 0)
	sortedList := sort.NewSortedListInt(&arrSlice)
	for i := 0; i < len(inputArr); i++ {
		sortedList.Append(inputArr[i])
	}
	fmt.Println(arrSlice)
}
