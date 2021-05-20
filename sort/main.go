package main

import (
	"fmt"
	"premetiv_sort/array_protector"
)

func main() {
	inputArr := [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}
	arrSlice := make([]int, 0)
	for i := 0; i < len(inputArr); i++ {
		arrSlice = array_protector.ProtectArrayQuickSort(arrSlice, inputArr[i])
	}
	fmt.Println(arrSlice)
}
